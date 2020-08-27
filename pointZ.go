package geo

import (
	"math"
)

type PointZ struct {
	X float64
	Y float64
	Z float64
}

func NewPointZ(x, y, z float64) *PointZ {
	return &PointZ{x, y, z}
}

func (p *PointZ) SetX(x float64) {
	p.X = x
}

func (p *PointZ) SetY(y float64) {
	p.Y = y
}

func (p *PointZ) SetZ(z float64) {
	p.Z = z
}

func EuclideanDisZ(p1, p2 PointZ) float64 {
	return math.Sqrt(p1.X*p1.X + p1.Y*p1.Y + p1.Z*p1.Z)
}

//判断元素是否相等
func (p1 PointZ) Equal(p2 PointZ) bool {
	return math.Abs(p1.X-p2.X) < COORDPRESION && math.Abs(p1.Y-p2.Y) < COORDPRESION && math.Abs(p1.Z-p2.Z) < COORDPRESION
}
