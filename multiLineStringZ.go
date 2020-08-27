package geo

type MultiLineStringZ []LineStringZ

func NewMultiLineStringZ(ls ...LineStringZ) *MultiLineStringZ {
	var ml MultiLineStringZ
	for _, v := range ls {
		ml = append(ml, v)
	}
	return &ml
}

func (l MultiLineStringZ) SetSrid(srid uint64) {
	SridMap[&l] = srid
}
