package geo

// https://stackoverflow.com/questions/563198/how-do-you-detect-where-two-line-segments-intersect
// 判断两个线段是否相交
func SegmentRelation(seg1, seg2 LineSegment) GeometryRealation {
	r := ConvertSegToVector(seg1)
	s := ConvertSegToVector(seg2)
	q_p := ConvertPointToVector(seg1.Start, seg2.Start)
	crss1 := r.Cross(*s)
	crss2 := q_p.Cross(*r)
	if crss1 == 0 && crss2 == 0 {
		return RELA_TOUCH
	} else if crss1 == 0 && crss2 != 0 {
		return RELA_DISJOINT
	} else if crss1 != 0 {
		//(q − p) × s / (r × s)
		res1 := q_p.Cross(*s) / (r.Cross(*s))
		res2 := q_p.Cross(*r) / (r.Cross(*s))
		if res1 > 0 && res1 < 1 && res2 > 0 && res2 < 1 {
			return RELA_INTERSECT
		}
		if (res1 == 0 || res1 == 1) && (res2 == 0 || res2 == 1) {
			return RELA_TOUCH
		}
	}
	return RELA_DISJOINT // No collision
}

func SegPolyRelation(seg LineSegment, poly Polygon) GeometryRealation {
	lr := poly.GetExteriorRing()
	pointCount := lr.GetPointCount()
	isPointOnBoundary := false
	for i := 0; i < pointCount-2; i++ {
		pntStart := lr[i]
		pntEnd := lr[i+1]
		relation := SegmentRelation(LineSegment{pntStart, pntEnd}, seg)
		if relation == RELA_INTERSECT {
			return RELA_INTERSECT
		} else if relation == RELA_TOUCH {
			isPointOnBoundary = true
		}
	}
	flag1 := PointInPolygon(seg.Start, poly)
	flag2 := PointInPolygon(seg.End, poly)
	//最后判断线段的首尾点在不在多边形内
	if flag1 && flag2 && !isPointOnBoundary {
		return RELA_CONTAIN
	} else if flag1 && flag2 && isPointOnBoundary {
		return RELA_COVER
	} else if !flag1 && !flag2 && isPointOnBoundary {
		return RELA_TOUCH
	} else if !flag1 && !flag2 && !isPointOnBoundary {
		return RELA_DISJOINT
	}
	return RELA_UNKNOWN
}

func LinePolyRelation(line LineString, poly Polygon) GeometryRealation {
	pointCount := line.GetPointCount()
	curRelation := RELA_UNKNOWN
	for i := 0; i < pointCount-1; i++ {
		Start := line[i]
		End := line[(i+1)%pointCount]
		relation := SegPolyRelation(LineSegment{Start, End}, poly)
		if relation == RELA_INTERSECT {
			return RELA_INTERSECT
		}
		if curRelation == RELA_UNKNOWN || curRelation == RELA_CONTAIN || curRelation == RELA_DISJOINT {
			curRelation = relation
		}
	}
	return curRelation
}

func LinearPolyRelation(line LinearRing, poly Polygon) GeometryRealation {
	//对环进行打断，去掉最后一个点
	pointCount := line.GetPointCount()
	curRelation := RELA_UNKNOWN
	for i := 0; i < pointCount-2; i++ {
		Start := line[i]
		End := line[i+1]
		relation := SegPolyRelation(LineSegment{Start, End}, poly)
		if relation == RELA_INTERSECT {
			return RELA_INTERSECT
		}
		if curRelation == RELA_UNKNOWN || curRelation == RELA_CONTAIN || curRelation == RELA_DISJOINT {
			curRelation = relation
		}
	}
	return curRelation
}

func PolyRelation(poly1, poly2 Polygon) GeometryRealation {
	lr := poly1.GetExteriorRing()
	return LinearPolyRelation(lr, poly2)
}

func IsPointOnLine(point Point, line LineString) bool {
	isOn := false
	for i := range line {
		if i == line.GetPointCount()-1 {
			isOn = IsPointOnSegment(line[0], line[line.GetPointCount()-1], point)
		} else {
			isOn = IsPointOnSegment(line[i], line[i+1], point)
		}
		if isOn {
			break
		}
	}
	return isOn
}

func IsPointOnSegment(p1, p2, point Point) bool {
	//保证Q点坐标在p1,p2之间 且叉积为0
	if (point.X-p1.X)*(p2.Y-p1.Y) == (p2.X-p1.X)*(point.Y-p1.Y) &&
		IsPointInBox(BoundingBox(NewLineString(p1, p2)), point) {
		return true
	}
	return false
}

func LineRelation(line1 LineString, line2 LineString) GeometryRealation {
	for i := 0; i < line1.GetPointCount()-1; i++ {
		for ii := 0; ii < line2.GetPointCount()-1; ii++ {
			relation := SegmentRelation(LineSegment{line1[i], line1[i+1]}, LineSegment{line2[ii], line2[ii+1]})
			if relation != RELA_DISJOINT {
				return RELA_INTERSECT
			}
		}
	}
	return RELA_DISJOINT
}
