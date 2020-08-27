package geo

type MultiPointZZ []PointZ

func NewMultiPointZ(pois ...PointZ) *MultiPointZ {
	var multiPointZ MultiPointZ
	for _, v := range pois {
		multiPointZ = append(multiPointZ, v)
	}
	return &multiPointZ
}

func (mp *MultiPointZ) Append(p PointZ) {
	*mp = append(*mp, p)
}

func (mp MultiPointZ) GetPointSet() []PointZ {
	return mp
}
