package geo

type MultiPoint []Point

func NewMultiPoint(pois ...Point) *MultiPoint {
	var multiPoint MultiPoint
	for _, v := range pois {
		multiPoint = append(multiPoint, v)
	}
	return &multiPoint
}

func (mp *MultiPoint) AddPoint(p Point) {
	*mp = append(*mp, p)
}

func (mp MultiPoint) SetSrid(srid uint64) {
	SridMap[&mp] = srid
}

func (mp MultiPoint) ToWkt() string {
	return PointToWkt(mp...)
}

func (mp MultiPoint) BoundingBox() Box {
	return calBox(mp...)
}

func (mp MultiPoint) TypeString() string {
	return "MultiPoint"
}
