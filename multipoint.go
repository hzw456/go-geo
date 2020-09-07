package geo

type MultiPoint []Point
type MultiPointZ []PointZ

func NewMultiPoint(pois ...Point) *MultiPoint {
	var multiPoint MultiPoint
	for _, v := range pois {
		multiPoint = append(multiPoint, v)
	}
	return &multiPoint
}

func (mp *MultiPoint) Append(p Point) {
	*mp = append(*mp, p)
}

func (mp MultiPoint) SetSrid(srid uint64) {
	SridMap[&mp] = srid
}

func (mp MultiPoint) GetPointSet() []Point {
	return mp
}
