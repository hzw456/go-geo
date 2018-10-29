package gogeo

const MAX = 9999999999999.0

type box struct {
	minX float64
	minY float64
	maxX float64
	maxY float64
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
	switch geo := geo.(type) {
	case Point:
		return calBox(geo)
	case MultiPoint, LineString, LinearRing:
		pois := geo.([]Point)
		return calBox(pois...)
	case MultiLineString:
		var pois []Point
		for _, v := range geo {
			for _, vv := range v {
				pois = append(pois, vv)
			}
		}
		return calBox(pois...)
	case Polygon:
		var pois []Point
		for _, v := range geo {
			for _, vv := range v {
				pois = append(pois, vv)
			}
		}
		return calBox(pois...)
	case MultiPolygon:
		var pois []Point
		for _, v := range geo {
			for _, vv := range v {
				for _, vvv := range vv {
					pois = append(pois, vvv)
				}
			}
		}
		return calBox(pois...)
	default:
		return calBox(Point{0, 0})
	}
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
