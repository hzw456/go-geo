package gogeo

import (
	"math"
)

const delta float64 = 1e-6

func dcmp(x float64) int {
	if math.Abs(x) < delta {
		return 0
	} else if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

//rebuild by hzw
func IsPointInPolygon(point Point, poly Polygon) bool {
	isIn := false
	lr := poly.GetExteriorRing()
	pointCount := lr.GetPointCount()
	for i := 0; i < pointCount-1; i++ {
		j := i + 1
		p1 := lr[i]
		p2 := lr[j]
		if IsPointOnSegment(p1, p2, point) {
			return true
		}
		//射线法判断 1 min(P1.y,P2.y)<P.y<=max(P1.y,P2.y) 2 point的射线相交 http://blog.letow.top/2017/11/13/vector-cross-product-cal-intersection/
		if (dcmp(p1.Y-point.Y) > 0) != (dcmp(p2.Y-point.Y) > 0) && dcmp(point.X-(point.Y-p1.Y)*(p1.X-p2.X)/(p1.Y-p2.Y)-p1.X) < 0 {
			isIn = !isIn
		}
	}
	return isIn
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

// void precalc_values() {

// 	int   i, j=polyCorners-1 ;

// 	for(i=0; i<polyCorners; i++) {
// 	  if(polyY[j]==polyY[i]) {
// 		constant[i]=polyX[i];
// 		multiple[i]=0; }
// 	  else {
// 		constant[i]=polyX[i]-(polyY[i]*polyX[j])/(polyY[j]-polyY[i])+(polyY[i]*polyX[i])/(polyY[j]-polyY[i]);
// 		multiple[i]=(polyX[j]-polyX[i])/(polyY[j]-polyY[i]); }
// 	  j=i; }}

//   bool pointInPolygon() {

// 	int   i, j=polyCorners-1 ;
// 	bool  oddNodes=NO      ;

// 	for (i=0; i<polyCorners; i++) {
// 	  if ((polyY[i]< y && polyY[j]>=y
// 	  ||   polyY[j]< y && polyY[i]>=y)) {
// 		oddNodes^=(y*multiple[i]+constant[i]<x); }
// 	  j=i; }

// 	return oddNodes; }
// //支持多点的快速判断

// //参照OGC标准的拓扑关系
// //包含关系，判断geom1是否包含geom2
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
