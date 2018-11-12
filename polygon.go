package gogeo

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
			if segmentIntersect(srcV0, srcV1, dstV0, dstV1) {
				return true
			}
		}
	}
	return false
}
