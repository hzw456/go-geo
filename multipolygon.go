package geo

type MultiPolygon []Polygon

func NewMultiPolygon(polys ...Polygon) *MultiPolygon {
	var mulitipoly MultiPolygon
	for _, v := range polys {
		mulitipoly = append(mulitipoly, v)
	}
	return &mulitipoly
}

func (mp MultiPolygon) SetSrid(srid uint64) {
	SridMap[&mp] = srid
}
