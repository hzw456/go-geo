package gogeo

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

//参照OGC标准的拓扑关系
//包含关系，判断geom1是否包含geom2
// func IsContains(geo1 Geometry, geo2 Geometry) {
// 	//求多边形的面积 论文:《多边形面积的计算与面积法的应用》
// 	var pois []Point
// 	switch geo := geo.(type) {
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
// 		return calBox(pois...)
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

func IsPointOnSegment(p1, p2, point Point) bool {
	//保证Q点坐标在p1,p2之间 且叉积为0
	if (point.X-p1.X)*(p2.Y-p1.Y) == (p2.X-p1.X)*(point.Y-p1.Y) &&
		IsPointInBox(Envelope(*NewLine(p1, p2)), point) {
		return true
	}
	return false
}
