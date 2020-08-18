package geo

import (
	"math"
)

// type Coord struct {
// 	lat float64
// 	lng float64
// }

type Point struct {
	X float64
	Y float64
}

type MultiPoint []Point

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func NewMultiPoint(pois ...Point) *MultiPoint {
	var multiPoint MultiPoint
	for _, v := range pois {
		multiPoint = append(multiPoint, v)
	}
	return &multiPoint
}

func (p *Point) SetX(x float64) {
	p.X = x
}

func (p *Point) SetY(y float64) {
	p.Y = y
}

func (multiPoint *MultiPoint) AddPoint(p Point) {
	*multiPoint = append(*multiPoint, p)
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

func (g Point) SetSRID(srid int) {
	SridMap[&g] = srid
}

func (g MultiPoint) SetSRID(srid int) {
	SridMap[&g] = srid
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
