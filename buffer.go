package geo

import (
	"math"

	"github.com/hzw456/go-geo/geos"
)

const CirclePolygonEdgeCount = 24

//点旋转法求buffer 论文：《一种GIS缓冲区矢量生成算法及实现》
func pointBuffer(p1 Point, bufferDis float64) Polygon {
	if bufferDis < 0 {
		return nil
	}
	bufferPoints := getBufferCoordsByRadian(p1, 0, 2*math.Pi, bufferDis)
	lr := NewLinearRing(bufferPoints...)
	poly := NewPolygon(*lr)
	return *poly
}

func getBufferCoordsByRadian(center Point, startRadian, endRadian, radius float64) []Point {
	gamma := 2 * math.Pi / CirclePolygonEdgeCount
	if radius < 0 {
		return nil
	}
	var bufferPoints []Point
	for phi := startRadian; phi <= endRadian; phi += gamma {
		point := NewPoint(center.X+radius*math.Cos(phi), center.Y+radius*math.Sin(phi))
		bufferPoints = append(bufferPoints, *point)
	}
	return bufferPoints
}

func lineBuffer(line1 LineString, bufferDis float64) Polygon {
	leftPoints := oneSideBuffer(line1, bufferDis)
	line2 := line1
	(&line2).Reverse()
	rightPoints := oneSideBuffer(line2, bufferDis)
	var bufferPoints []Point
	bufferPoints = append(bufferPoints, leftPoints...)
	bufferPoints = append(bufferPoints, rightPoints...)
	lr := NewLinearRing(bufferPoints...)
	poly := NewPolygon(*lr)
	return *poly
}

func oneSideBuffer(line1 LineString, bufferDis float64) []Point {
	err := line1.Verify()
	if err != nil {
		return nil
	}
	var bufferPoints []Point
	//第一节点的缓冲区
	quadrantAngle, err := CacQuadrantAngle(line1[0], line1[1])
	if err != nil {
		return nil
	}
	startRadian := quadrantAngle + math.Pi
	endRadian := quadrantAngle + (3*math.Pi)/2
	bufferPoints = append(bufferPoints, getBufferCoordsByRadian(line1[0], startRadian, endRadian, bufferDis)...)

	//中间节点缓冲区
	for i := 1; i < len(line1)-1; i++ {
		convex, err := CacConvex(line1[i-1], line1[i], line1[i+1])
		if err != nil {
			return nil
		}
		//如果是平角，则忽略这个点
		if convex == 2 {
			continue
		}
		quadrantAngle, err = CacQuadrantAngle(line1[i], line1[i-1])
		if err != nil {
			return nil
		}
		angle, err := CacAngle(line1[i-1], line1[i], line1[i+1])
		if err != nil {
			return nil
		}
		if convex == 0 {
			startRadian = quadrantAngle + math.Pi/2
			endRadian = quadrantAngle - math.Pi/2 + angle
			bufferPoints = append(bufferPoints, getBufferCoordsByRadian(line1[i], startRadian, endRadian, bufferDis)...)
		} else if convex == 1 {
			phi := quadrantAngle + angle/2
			point := NewPoint(line1[i].X+bufferDis*math.Cos(phi), line1[i].Y+bufferDis*math.Sin(phi))
			bufferPoints = append(bufferPoints, *point)
		}
	}
	//最后一个点的缓冲区
	quadrantAngle, err = CacQuadrantAngle(line1[line1.GetPointCount()-2], line1.GetEndPoint())
	startRadian = quadrantAngle + (3*math.Pi)/2
	endRadian = quadrantAngle + 2*math.Pi
	bufferPoints = append(bufferPoints, getBufferCoordsByRadian(line1.GetEndPoint(), startRadian, endRadian, bufferDis)...)
	return bufferPoints
}

func polyBuffer(p1 Polygon, width float64) Polygon {
	var pois []geos.Point
	for _, poi := range p1.GetExteriorPoints() {
		pois = append(pois, geos.Point{poi.X, poi.Y})
	}
	cg := geos.CreatePolygon(pois)
	geosm := cg.Buffer(width)
	geom := GeosToGeo(geosm)
	if geom != nil && geom.Type() == GEOMETRY_POLYGON {
		return geom.(Polygon)
	}
	return Polygon{}
}
