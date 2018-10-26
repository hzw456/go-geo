package gogeo

import (
	"math"
)

type Point struct {
	X, Y float64
}

type MultiPoint []Point

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
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
