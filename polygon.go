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
