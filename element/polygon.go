package element

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
