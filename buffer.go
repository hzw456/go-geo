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

func bufferWithStyle(g LineString, width float64, quadsegs int, endCapStyle geos.CapType, joinStyle geos.JoinType, mitreLimit float64) Polygon {
	var pois []geos.Point
	for _, poi := range g {
		pois = append(pois, geos.Point{poi.X, poi.Y})
	}
	cg := geos.CreateLineString(pois)
	geosm := cg.BufferWithStyle(width, quadsegs, endCapStyle, joinStyle, mitreLimit)
	geom := GeosToGeo(geosm)
	if geom != nil && geom.Type() == GEOMETRY_POLYGON {
		return geom.(Polygon)
	}
	return Polygon{}
}
