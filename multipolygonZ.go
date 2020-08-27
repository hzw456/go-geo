package geo

type MultiPolygonZ []PolygonZ

func NewMultiPolygonZ(polys ...PolygonZ) *MultiPolygonZ {
	var mulitipoly MultiPolygonZ
	for _, v := range polys {
		mulitipoly = append(mulitipoly, v)
	}
	return &mulitipoly
}
