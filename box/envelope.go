package box

import (
	"github.com/sadnessly/go-geo/element"
)

const MAX = 9999999999999.0

//求多边形的面积 论文:《多边形面积的计算与面积法的应用》
func Envelope(geo element.Geometry) box {
	switch geo := geo.(type) {
	case element.Point:
		return calBox(geo)
	case element.MultiPoint, element.LineString, element.LinearRing:
		pois := geo.([]element.Point)
		return calBox(pois...)
	case element.MultiLineString:
		var pois []element.Point
		for _, v := range geo {
			for _, vv := range v {
				pois = append(pois, vv)
			}
		}
		return calBox(pois...)
	default:
		return calBox(element.Point{0, 0})
	}
}

func calBox(points ...element.Point) box {
	var minX, minY, maxX, maxY float64 = MAX, MAX, -MAX, -MAX
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
	return box{minX, minY, maxX, maxY}
}
