package geo

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) SetX(x float64) {
	p.X = x
}

func (p *Point) SetY(y float64) {
	p.Y = y
}

func (g Point) SetSrid(srid uint64) {
	SridMap[&g] = srid
}

func (p Point) BoundingBox() Box {
	return calBox(p)
}

func (p Point) Buffer(width float64) Polygon {
	return pointBuffer(p, width)
}

//Euclidean distance
func (p1 Point) Distance(p2 Point) float64 {
	if SRID_WGS84_GPS == SRID_WGS84_GPS {
		return CoordDistance(p1, p2)
	} else {
		return EuclideanDis(p1, p2)
	}
}

func EuclideanDis(p1, p2 Point) float64 {
	return math.Sqrt(p1.X*p1.X + p1.Y*p1.Y)
}

//判断元素是否相等
func (p1 Point) Equal(p2 Point) bool {
	return math.Abs(p1.X-p2.X) < COORDPRESION && math.Abs(p1.Y-p2.Y) < COORDPRESION
}

func CoordDistance(p1, p2 Point) float64 {
	lat1 := p1.X * math.Pi / 180.0
	lng1 := p1.Y * math.Pi / 180.0
	lat2 := p2.X * math.Pi / 180.0
	lng2 := p2.Y * math.Pi / 180.0
	return math.Acos(math.Sin(lat1)*math.Sin(lat2)+math.Cos(lat1)*math.Cos(lat2)*math.Cos(lng2-lng1)) * EARTHRADIUSMI
}

func CoordHaversine(p1, p2 Point) float64 {
	lat1 := p1.X * math.Pi / 180.0
	lng1 := p1.Y * math.Pi / 180.0
	lat2 := p2.X * math.Pi / 180.0
	lng2 := p2.Y * math.Pi / 180.0
	a := math.Pow(math.Sin((lat2-lat1)/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin((lng2-lng1)/2), 2)
	return 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a)) * EARTHRADIUSMI
}
