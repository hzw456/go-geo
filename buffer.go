package geo

import (
	"math"
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
