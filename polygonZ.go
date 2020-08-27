package geo

type PolygonZ []LinearRingZ

func NewPolygonZ(lrs ...LinearRingZ) *PolygonZ {
	var rings []LinearRingZ
	for _, v := range lrs {
		rings = append(rings, v)
	}
	poly := PolygonZ(rings)
	return &poly
}

func NewPolygonZFromPois(pts ...PointZ) *PolygonZ {
	ring := NewLinearRingZ(pts...)
	return NewPolygonZ(*ring)
}

//多边形的外环
func (poly PolygonZ) GetExteriorRing() LinearRingZ {
	if len(poly) == 0 {
		return nil
	}
	return poly[0]
}

func (poly *PolygonZ) SetExteriorRing(line LinearRingZ) {
	if len(*poly) == 0 {
		*poly = append(*poly, line)
	} else {
		(*poly)[0] = line
	}
}

func (poly PolygonZ) GetInteriorRing() []LinearRingZ {
	if len(poly) == 0 {
		return nil
	}
	return poly[1:]
}

//多边形内的洞
func (poly *PolygonZ) AddInteriorRing(lines LinearRingZ) {
	*poly = append(*poly, lines)
}

func (multipoly *MultiPolygonZ) AddPolygonZ(poly PolygonZ) {
	*multipoly = append(*multipoly, poly)
}

func (poly PolygonZ) GetExteriorPoints() []PointZ {
	return poly.GetExteriorRing()
}
