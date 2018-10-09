package element

import (
	"math"
)

const CirclePolygonEdgeCount = 8

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func NewPointStr(pointStr string) {

}

func NewPointJson(pointJson string) {

}

func (p *Point) SetX(x float64) {
	p.X = x
}

func (p *Point) SetY(y float64) {
	p.Y = y
}

//Euclidean distance
func (p1 Point) PointDistance(p2 Point) float64 {
	return math.Sqrt(p1.X*p1.X + p1.Y*p1.Y)
}

//判断元素是否相等
func (p1 Point) Equal(p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

//点旋转法求buffer 论文：《一种GIS缓冲区矢量生成算法及实现》
func (p1 Point) Buffer(bufferDis float64) Polygon {
	var bufferPoints []Point
	for i := 0; i < CirclePolygonEdgeCount; i++ {
		point := NewPoint(0, 0)
		point.SetX(p1.X + bufferDis*math.Sin(2*math.Pi/CirclePolygonEdgeCount*float64(i)))
		point.SetY(p1.Y + bufferDis*math.Cos(2*math.Pi/CirclePolygonEdgeCount*float64(i)))
		bufferPoints = append(bufferPoints, *point)
	}
	lr := NewLinearRing(*NewLine2(bufferPoints))
	poly := NewPolygon(*lr)
	return poly
}
