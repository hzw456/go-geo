package geo

import (
	"math"
)

type Point struct {
	//如果用经纬度，x为经度，y为纬度
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

func (p Point) Buffer(width float64) Polygon {
	return pointBuffer(p, width)
}

//Euclidean distance
func (p1 Point) Distance(p2 Point) float64 {
	return EuclideanDis(p1, p2)
}

func EuclideanDis(p1, p2 Point) float64 {
	return math.Sqrt((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))
}

//判断元素是否相等
func (p1 Point) Equal(p2 Point) bool {
	return math.Abs(p1.X-p2.X) < COORDPRESION && math.Abs(p1.Y-p2.Y) < COORDPRESION
}

//HaverSin
func CoordDistance(p1, p2 Point) float64 {
	lat1 := p1.Y * math.Pi / 180.0
	lng1 := p1.X * math.Pi / 180.0
	lat2 := p2.Y * math.Pi / 180.0
	lng2 := p2.X * math.Pi / 180.0
	a := math.Pow(math.Sin((lat2-lat1)/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin((lng2-lng1)/2), 2)
	return 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a)) * EARTHRADIUSMI
}

func CoordGreateCircle(p1, p2 Point) float64 {
	lat1 := p1.Y * math.Pi / 180.0
	lng1 := p1.X * math.Pi / 180.0
	lat2 := p2.Y * math.Pi / 180.0
	lng2 := p2.X * math.Pi / 180.0
	return math.Acos(math.Sin(lat1)*math.Sin(lat2)+math.Cos(lat1)*math.Cos(lat2)*math.Cos(lng2-lng1)) * EARTHRADIUSMI
}

// 根据距离和方位角进行下一个点的计算，TODO
func DestinationPoint(lat, lon, distance, bearing float64) (float64, float64) {

	var toRadians = func(v float64) float64 { return v * math.Pi / 180 }
	var toDegrees = func(v float64) float64 { return v * 180 / math.Pi }

	// sinφ2 = sinφ1·cosδ + cosφ1·sinδ·cosθ
	// tanΔλ = sinθ·sinδ·cosφ1 / cosδ−sinφ1·sinφ2
	// see mathforum.org/library/drmath/view/52049.html for derivation

	var δ = distance / EARTHRADIUSMI // angular distance in radians
	var θ = toRadians(bearing)

	var φ1 = toRadians(lat)
	var λ1 = toRadians(lon)

	sinφ2 := math.Sin(φ1)*math.Cos(δ) + math.Cos(φ1)*math.Sin(δ)*math.Cos(θ)
	φ2 := math.Asin(sinφ2)
	y := math.Sin(θ) * math.Sin(δ) * math.Cos(φ1)
	x := math.Cos(δ) - math.Sin(φ1)*math.Sin(φ2)
	λ2 := λ1 + math.Atan2(y, x)
	return toDegrees(φ2), math.Mod((toDegrees(λ2)+540), 360) - 180 // normalise to −180..+180°
}
