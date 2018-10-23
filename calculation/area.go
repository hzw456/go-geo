package calculation

import (
	"math"

	"github.com/sadnessly/go-geo/element"
)

//求多边形的面积 论文:《多边形面积的计算与面积法的应用》
func Area(geo element.Geometry) float64 {
	switch geo := geo.(type) {
	case element.Polygon:
		return polyArea(geo)
	case element.MultiPolygon:
		return MultiPolyArea(geo)
	}
	return 0
}

func polyArea(poly element.Polygon) float64 {
	lr := poly.GetExteriorRing()
	if lr == nil {
		return 0
	}
	ptCount := lr.GetPointCount() - 1
	var area float64
	for i := 0; i < ptCount; i++ {
		//最后一个点的处理
		j := (i + 1) % ptCount
		area += lr[i].X * lr[j].Y
		area -= lr[i].Y * lr[j].X
	}
	area /= 2
	return math.Abs(area)
}

func MultiPolyArea(multiPoly element.MultiPolygon) float64 {
	areaSum := 0.0
	for _, v := range multiPoly {
		areaSum += polyArea(v)
	}
	return areaSum
}
