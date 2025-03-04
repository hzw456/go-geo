package geo

import "math"

// const delta float64 = 1e-6

func dcmp(x float64) int {
	if math.Abs(x) < COORDPRESION {
		return 0
	} else if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

// rebuild by hzw
func PointInPolygon(point Point, poly Polygon) bool {
	isIn := false
	lr := poly.GetExteriorRing()
	pointCount := lr.GetPointCount()
	for i := 0; i < pointCount-1; i++ {
		j := i + 1
		p1 := lr[i]
		p2 := lr[j]
		//在边界上的判断
		if IsPointOnSegment(p1, p2, point) {
			return true
		}
		//射线法判断 1 min(P1.y,P2.y)<P.y<=max(P1.y,P2.y) 2 point的射线相交 http://blog.letow.top/2017/11/13/vector-cross-product-cal-intersection/
		if (dcmp(p1.Y-point.Y) > 0) != (dcmp(p2.Y-point.Y) > 0) && dcmp(point.X-(point.Y-p1.Y)*(p1.X-p2.X)/(p1.Y-p2.Y)-p1.X) < 0 {
			isIn = !isIn
		}
	}
	if isIn {
		return true
	} else {
		return false
	}
}
