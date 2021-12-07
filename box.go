package geo

import (
	"math"
)

type Box struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

func NewBox(minX, minY, maxX, maxY float64) *Box {
	return &Box{MinX: minX, MinY: minY, MaxX: maxX, MaxY: maxY}
}

func NewRectBox(p Point, lengthX, lengthY float64) (r *Box) {
	return &Box{MinX: p.X - lengthX, MinY: p.Y - lengthY, MaxX: p.X + lengthX, MaxY: p.Y + lengthY}
}

//在边界上也算在内部
func IsPointInBox(b *Box, p Point) bool {
	if b.MinX <= p.X && p.X <= b.MaxX && b.MinY <= p.Y && p.Y <= b.MaxY {
		return true
	}
	return false
}

func BoxToGeo(b Box) Geometry {
	p1 := Point{b.MinX, b.MinY}
	p2 := Point{b.MinX, b.MaxY}
	p3 := Point{b.MaxX, b.MaxY}
	p4 := Point{b.MaxX, b.MinY}

	if p1.Equal(p3) {
		//元素是个点
		return p1
	} else if p1.Equal(p2) {
		//元素是条线 y坐标不同
		return LineString{p1, p3}
	} else if p2.Equal(p3) {
		//元素是条线 x坐标不同
		return LineString{p1, p2}
	}
	return *NewPolygon(LinearRing{p1, p2, p3, p4})
}

func calBox(points ...Point) *Box {
	var minX, minY, maxX, maxY float64 = INF, INF, -INF, -INF
	for _, v := range points {
		if minX > v.X {
			minX = v.X
		}
		if minY > v.Y {
			minY = v.Y
		}
		if maxX < v.X {
			maxX = v.X
		}
		if maxY < v.Y {
			maxY = v.Y
		}
	}
	return &Box{minX, minY, maxX, maxY}
}

func BoundingBox(geom Geometry) *Box {
	switch geom := geom.(type) {
	case Point:
		return calBox(geom)
	case MultiPoint:
		return calBox(geom...)
	case LineString:
		return calBox(geom...)
	case MultiLineString:
		var pois []Point
		for _, v := range geom {
			for _, vv := range v {
				pois = append(pois, vv)
			}
		}
		return calBox(pois...)
	case Polygon:
		return calBox(geom.GetExteriorPoints()...)
	case MultiPolygon:
		var pois []Point
		for _, v := range geom {
			for _, vv := range v {
				for _, vvv := range vv {
					pois = append(pois, vvv)
				}
			}
		}
		return calBox(pois...)
	}
	return &Box{}
}

func (b1 *Box) Intersect(b2 *Box) bool {
	if b1 == nil || b2 == nil {
		return false
	}
	if b1.MaxX < b2.MinX || b1.MinX > b2.MaxX || b1.MinY > b2.MaxY || b1.MaxY < b2.MinY {
		return false
	}
	return true
}

func (b1 *Box) Union(b2 *Box) *Box {
	// math.Max(float64, float64) float64
	return &Box{MinX: math.Min(b1.MinX, b2.MinX), MinY: math.Min(b1.MinY, b2.MinY),
		MaxX: math.Max(b1.MaxX, b2.MaxX), MaxY: math.Max(b1.MinY, b2.MinY)}
}

func BoxUnion(boxs ...*Box) *Box {
	if len(boxs) == 0 {
		return &Box{}
	}
	box := boxs[0]
	for _, b1 := range boxs[1:] {
		box = box.Union(b1)
	}
	return box
}

// Size computes the measure of a rectangle (the product of its side lengths).
func (b1 *Box) Size() float64 {
	return (b1.MaxX - b1.MinX) * (b1.MaxY - b1.MinY)
}

func (b1 *Box) Contain(b2 *Box) bool {
	if b1.MinX > b2.MinX || b2.MaxX > b1.MaxX {
		return false
	}
	if b1.MinY > b2.MinY || b2.MaxY > b1.MaxY {
		return false
	}
	return true
}
