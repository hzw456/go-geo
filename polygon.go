package gogeo

import "errors"

type Polygon []LinearRing
type MultiPolygon []Polygon

func NewPolygon(lrs ...LinearRing) *Polygon {
	var rings []LinearRing
	for _, v := range lrs {
		rings = append(rings, v)
	}
	poly := Polygon(rings)
	return &poly
}

func NewSimplePolygon(pts ...Point) *Polygon {
	var ring LinearRing
	for _, v := range pts {
		ring = append(ring, v)
	}
	var rings []LinearRing
	rings = append(rings, ring)
	poly := Polygon(rings)
	return &poly
}

func NewMultiPolygon(polys ...Polygon) *MultiPolygon {
	var mulitipoly MultiPolygon
	for _, v := range polys {
		mulitipoly = append(mulitipoly, v)
	}
	return &mulitipoly
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
func (poly *Polygon) AddInteriorRing(lines LinearRing) {
	*poly = append(*poly, lines)
}

func (multipoly *MultiPolygon) AddPolygon(poly Polygon) {
	*multipoly = append(*multipoly, poly)
}

func (p *Polygon) Verify() error {
	extring := p.GetExteriorRing()
	ptCount := extring.GetPointCount() - 1
	if ptCount < 3 {
		return errors.New("polygon invaild, point less than 3")
	}
	for i := 0; i < ptCount; i++ {
		if extring[i].X > 180.0 || extring[i].X < -180.0 || extring[i].Y > 90.0 || extring[i].Y < -90.0 {
			return errors.New("lnglat invaild, lat[-90,90] lng[-180,180]")
		}
	}
	//每次计算面积判断是否合法会有点复杂，考虑采用其他的方式判断
	// if p.Area() < 0.0000000001 {
	// 	return errors.New("polygon invaild, area equal 0")
	// }
	if p.SelfIntersect() {
		return errors.New("polygon self-intersect")
	}
	return nil
}

func (poly Polygon) SelfIntersect() bool {
	//if edgeNum > 3 {
	//	if p.Points[edgeNum-2].Equal(&p.Points[edgeNum-1]) || p.Points[0].Equal(&p.Points[edgeNum-1]) {
	//		edgeNum = edgeNum - 1
	//	}
	//}
	exRing := poly.GetExteriorRing()
	pointCount := exRing.GetPointCount()
	for i := 0; i < pointCount-1; i++ {
		srcV0 := exRing[i]
		srcV1 := exRing[(i+1)%pointCount]
		for j := 0; j < pointCount; j++ {
			if i == j || i-j == 1 || j-i == 1 || i-j == pointCount-1 || j-i == pointCount-1 { //cojoin or ewqul
				continue
			}
			dstV0 := exRing[j]
			dstV1 := exRing[(j+1)%pointCount]
			if segmentIntersect(LineSegment{srcV0, srcV1}, LineSegment{dstV0, dstV1}) {
				return true
			}
		}
	}
	return false
}
