package gogeo

func segmentIntersect(seg1, seg2 LineSegment) bool {
	s1X := seg1.end.X - seg1.start.X
	s1Y := seg1.end.Y - seg1.start.Y
	s2X := seg2.end.X - seg2.start.X
	s2Y := seg2.end.Y - seg2.start.Y

	s := (-s1Y*(seg1.start.X-seg2.start.X) + s1X*(seg1.start.Y-seg2.start.Y)) / (-s2X*s1Y + s1X*s2Y)
	t := (s2X*(seg1.start.Y-seg2.start.Y) - s2Y*(seg1.start.X-seg2.start.X)) / (-s2X*s1Y + s1X*s2Y)

	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
		return true
	}
	return false // No collision
}

func SegPolyRealation(seg LineSegment, poly Polygon) int {
	lr := poly.GetExteriorRing()
	pointCount := lr.GetPointCount()
	for i := 0; i < pointCount-1; i++ {
		pntStart := lr[i]
		pntEnd := lr[(i+1)%pointCount]
		if segmentIntersect(LineSegment{pntStart, pntEnd}, seg) {
			return GEO_INTERSECT
		}
	}
	flag1 := IsPointInPolygon(seg.start, poly)
	flag2 := IsPointInPolygon(seg.end, poly)
	if flag1 && flag2 {
		return GEO_CONTAIN
	}
	if !flag1 && !flag2 {
		return GEO_DISJOINT
	}
	return GEO_UNKNOWN
}

// func IsIntersect(geo1, geo2 Geometry) bool {
// 	var pois []Point
// 	switch geo := geo1.(type) {
// 	case Point:
// 		return calBox(geo)
// 	case MultiPoint:
// 		for _, v := range geo {
// 			pois = append(pois, v)
// 		}
// 	case LineString:
// 		for _, v := range geo {
// 			pois = append(pois, v)
// 		}
// 	case LinearRing:
// 		for _, v := range geo {
// 			pois = append(pois, v)
// 		}
// 	case MultiLineString:
// 		for _, v := range geo {
// 			for _, vv := range v {
// 				pois = append(pois, vv)
// 			}
// 		}
// 	case Polygon:
// 		for _, v := range geo {
// 			for _, vv := range v {
// 				pois = append(pois, vv)
// 			}
// 		}
// 	case MultiPolygon:
// 		for _, v := range geo {
// 			for _, vv := range v {
// 				for _, vvv := range vv {
// 					pois = append(pois, vvv)
// 				}
// 			}
// 		}
// 	default:
// 		return calBox(Point{0, 0})
// 	}
// 	return calBox(pois...)
// }
