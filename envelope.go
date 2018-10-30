package gogeo

const MAX = 9999999999999.0

type box struct {
	minX float64
	minY float64
	maxX float64
	maxY float64
}

//在边界上也算在内部
func IsPointInBox(b box, p Point) bool {
	if b.minX <= p.X && p.X <= b.maxX && b.minY <= p.Y && p.Y <= b.maxY {
		return true
	}
	return false
}

func BoxToGeo(b box) Geometry {
	p1 := Point{b.minX, b.minY}
	p2 := Point{b.minX, b.maxY}
	p3 := Point{b.maxX, b.maxY}
	p4 := Point{b.maxX, b.minY}

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

//求多边形的面积 论文:《多边形面积的计算与面积法的应用》
func Envelope(geo Geometry) box {
	var pois []Point
	switch geo := geo.(type) {
	case Point:
		return calBox(geo)
	case MultiPoint:
		for _, v := range geo {
			pois = append(pois, v)
		}
	case LineString:
		for _, v := range geo {
			pois = append(pois, v)
		}
	case LinearRing:
		for _, v := range geo {
			pois = append(pois, v)
		}
	case MultiLineString:
		for _, v := range geo {
			for _, vv := range v {
				pois = append(pois, vv)
			}
		}
	case Polygon:
		for _, v := range geo {
			for _, vv := range v {
				pois = append(pois, vv)
			}
		}
		return calBox(pois...)
	case MultiPolygon:
		for _, v := range geo {
			for _, vv := range v {
				for _, vvv := range vv {
					pois = append(pois, vvv)
				}
			}
		}
	default:
		return calBox(Point{0, 0})
	}
	return calBox(pois...)
}

func calBox(points ...Point) box {
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
