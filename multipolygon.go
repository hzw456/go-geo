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

func (mp MultiPolygon) ToWkt() string {
	return PolygonToWkt(mp...)
}

func (mp MultiPolygon) TypeString() string {
	return "MultiPolygon"
}

func (mp MultiPolygon) BoundingBox() Box {
	var pois []Point
	for _, v := range mp {
		for _, vv := range v {
			for _, vvv := range vv {
				pois = append(pois, vvv)
			}
		}
	}
	return calBox(pois...)
}
