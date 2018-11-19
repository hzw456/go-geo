package gogeo

//https://stackoverflow.com/questions/563198/how-do-you-detect-where-two-line-segments-intersect
//判断两个线段是否相交
func SegmentRelation(seg1, seg2 LineSegment) int {
	r := ConvertSegToVector(seg1)
	s := ConvertSegToVector(seg2)
	q_p := ConvertPointToVector(seg1.Start, seg2.Start)
	crss1 := r.Cross(*s)
	crss2 := q_p.Cross(*r)
	if crss1 == 0 && crss2 == 0 {
		return GEO_TOUCH
	} else if crss1 == 0 && crss2 != 0 {
		return GEO_DISJOINT
	} else if crss1 != 0 {
		//(q − p) × s / (r × s)
		res1 := q_p.Cross(*s) / (r.Cross(*s))
		res2 := q_p.Cross(*r) / (r.Cross(*s))
		if res1 > 0 && res1 < 1 && res2 > 0 && res2 < 1 {
			return GEO_INTERSECT
		}
		if (res1 == 0 || res1 == 1) && (res2 == 0 || res2 == 1) {
			return GEO_TOUCH
		}
	}
	return GEO_DISJOINT // No collision
}

// func Intersect(p0_x, p0_y, p1_x, p1_y, p2_x, p2_y, p3_x, p3_y float64) bool {
// 	s1_x := p1_x - p0_x
// 	s1_y := p1_y - p0_y
// 	s2_x := p3_x - p2_x
// 	s2_y := p3_y - p2_y

// 	s := (-s1_y*(p0_x-p2_x) + s1_x*(p0_y-p2_y)) / (-s2_x*s1_y + s1_x*s2_y)
// 	t := (s2_x*(p0_y-p2_y) - s2_y*(p0_x-p2_x)) / (-s2_x*s1_y + s1_x*s2_y)
// 	fmt.Println("res1=", s)
// 	fmt.Println("res2=", t)
// 	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
// 		// Collision detected
// 		//                if (i_x != NULL)
// 		//                    *i_x = p0_x + (t * s1_x);
// 		//                if (i_y != NULL)
// 		//                    *i_y = p0_y + (t * s1_y);
// 		return true
// 	}
// 	return false // No collision
// }

func SegPolyRelation(seg LineSegment, poly Polygon) int {
	lr := poly.GetExteriorRing()
	pointCount := lr.GetPointCount()
	isPointOnBoundary := false
	for i := 0; i < pointCount-1; i++ {
		pntStart := lr[i]
		pntEnd := lr[(i+1)%pointCount]
		relation := SegmentRelation(LineSegment{pntStart, pntEnd}, seg)
		if relation == GEO_INTERSECT {
			return GEO_INTERSECT
		} else if relation == GEO_TOUCH {
			isPointOnBoundary = true
		}
	}
	flag1 := IsPointInPolygon(seg.Start, poly) == GEO_CONTAIN || IsPointInPolygon(seg.Start, poly) == GEO_TOUCH
	flag2 := IsPointInPolygon(seg.End, poly) == GEO_CONTAIN || IsPointInPolygon(seg.End, poly) == GEO_TOUCH
	//最后判断线段的首尾点在不在多边形内
	if flag1 && flag2 && !isPointOnBoundary {
		return GEO_CONTAIN
	} else if flag1 && flag2 && isPointOnBoundary {
		return GEO_COVER
	} else if !flag1 && !flag2 && isPointOnBoundary {
		return GEO_TOUCH
	} else if !flag1 && !flag2 && !isPointOnBoundary {
		return GEO_DISJOINT
	}
	return GEO_UNKNOWN
}

func PolyRelation(poly1, poly2 Polygon) int {
	lr := poly1.GetExteriorRing()
	pointCount := lr.GetPointCount()
	curRelation := GEO_UNKNOWN
	for i := 0; i < pointCount-1; i++ {
		Start := lr[i]
		End := lr[(i+1)%pointCount]
		relation := SegPolyRelation(LineSegment{Start, End}, poly2)
		if relation == GEO_INTERSECT {
			return GEO_INTERSECT
		}
		if curRelation == GEO_UNKNOWN || curRelation == GEO_CONTAIN || curRelation == GEO_DISJOINT {
			curRelation = relation
		}
	}
	return curRelation
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
