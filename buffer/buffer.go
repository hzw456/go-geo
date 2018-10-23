package buffer

import (
	"math"

	"github.com/sadnessly/go-geo/calculation"
	"github.com/sadnessly/go-geo/element"
)

const CirclePolygonEdgeCount = 24

//点旋转法求buffer 论文：《一种GIS缓冲区矢量生成算法及实现》
func Buffer(geometry interface{}, bufferDis float64) element.Polygon {
	switch geometry.(type) {
	case element.Point:
		return pointBuffer(geometry.(element.Point), bufferDis)
	case element.LineString:
		return lineBuffer(geometry.(element.LineString), bufferDis)
		// case element.Polygon:
		// 	return polyBuffer(geometry.(element.Polygon), bufferDis)
	}
	return nil
}

func pointBuffer(p1 element.Point, bufferDis float64) element.Polygon {
	if bufferDis < 0 {
		return nil
	}
	bufferPoints := getBufferCoordsByRadian(p1, 0, 2*math.Pi, bufferDis)
	lr := element.NewLinearRing(*element.NewLine(bufferPoints...))
	poly := element.NewPolygon(*lr)
	return *poly
}

func getBufferCoordsByRadian(center element.Point, startRadian, endRadian, radius float64) []element.Point {
	gamma := 2 * math.Pi / CirclePolygonEdgeCount
	if radius < 0 {
		return nil
	}
	var bufferPoints []element.Point
	for phi := startRadian; phi <= endRadian; phi += gamma {
		point := element.NewPoint(center.X+radius*math.Cos(phi), center.Y+radius*math.Sin(phi))
		bufferPoints = append(bufferPoints, *point)
	}
	return bufferPoints
}

func lineBuffer(line1 element.LineString, bufferDis float64) element.Polygon {
	leftPoints := oneSideBuffer(line1, bufferDis)
	line2 := line1
	(&line2).Reverse()
	rightPoints := oneSideBuffer(line2, bufferDis)
	var bufferPoints []element.Point
	bufferPoints = append(bufferPoints, leftPoints...)
	bufferPoints = append(bufferPoints, rightPoints...)
	lr := element.NewLinearRing(*element.NewLine(bufferPoints...))
	poly := element.NewPolygon(*lr)
	return *poly
}

func oneSideBuffer(line1 element.LineString, bufferDis float64) []element.Point {
	err := line1.Verify()
	if err != nil {
		return nil
	}
	var bufferPoints []element.Point
	//第一节点的缓冲区
	quadrantAngle, err := calculation.CacQuadrantAngle(line1[0], line1[1])
	if err != nil {
		return nil
	}
	startRadian := quadrantAngle + math.Pi
	endRadian := quadrantAngle + (3*math.Pi)/2
	bufferPoints = append(bufferPoints, getBufferCoordsByRadian(line1[0], startRadian, endRadian, bufferDis)...)

	//中间节点缓冲区
	for i := 1; i < len(line1)-1; i++ {
		convex, err := calculation.CacConvex(line1[i-1], line1[i], line1[i+1])
		if err != nil {
			return nil
		}
		//如果是平角，则忽略这个点
		if convex == 2 {
			continue
		}
		quadrantAngle, err = calculation.CacQuadrantAngle(line1[i], line1[i-1])
		if err != nil {
			return nil
		}
		angle, err := calculation.CacAngle(line1[i-1], line1[i], line1[i+1])
		if err != nil {
			return nil
		}
		if convex == 0 {
			startRadian = quadrantAngle + math.Pi/2
			endRadian = quadrantAngle - math.Pi/2 + angle
			bufferPoints = append(bufferPoints, getBufferCoordsByRadian(line1[i], startRadian, endRadian, bufferDis)...)
		} else if convex == 1 {
			phi := quadrantAngle + angle/2
			point := element.NewPoint(line1[i].X+bufferDis*math.Cos(phi), line1[i].Y+bufferDis*math.Sin(phi))
			bufferPoints = append(bufferPoints, *point)
		}
	}
	//最后一个点的缓冲区
	quadrantAngle, err = calculation.CacQuadrantAngle(line1[line1.GetPointCount()-2], line1.GetEndPoint())
	startRadian = quadrantAngle + (3*math.Pi)/2
	endRadian = quadrantAngle + 2*math.Pi
	bufferPoints = append(bufferPoints, getBufferCoordsByRadian(line1.GetEndPoint(), startRadian, endRadian, bufferDis)...)
	return bufferPoints
}
