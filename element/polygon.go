package element

import (
	"math"
)

type Polygon []LinearRing

func NewPolygon(lr LinearRing) Polygon {
	var poly Polygon
	poly.SetExteriorRing(lr)
	return poly
}

//多边形的外环
func (poly Polygon) GetExteriorRing() LinearRing {
	if len(poly) == 0 {
		return nil
	}
	return poly[0]
}

func (poly *Polygon) SetExteriorRing(line LinearRing) {
	if len(*poly) == 0 {
		*poly = append(*poly, line)
	} else {
		(*poly)[0] = line
	}
}

func (poly Polygon) GetInteriorRing() []LinearRing {
	if len(poly) == 0 {
		return nil
	}
	return poly[1:]
}

//多边形内的洞
func (poly *Polygon) SetInteriorRing(lines LinearRing) {
	*poly = append(*poly, lines)
}

//求多边形的面积 论文:《多边形面积的计算与面积法的应用》
func (poly *Polygon) Area() float64 {
	lr := poly.GetExteriorRing()
	if lr == nil {
		return 0
	}
	ptCount := lr.GetPointCount()
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
