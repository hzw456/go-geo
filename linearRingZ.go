package geo

type LinearRingZ []PointZ

func NewLinearRingZ(point ...PointZ) *LinearRingZ {
	var points []PointZ
	for _, point := range point {
		points = append(points, point)
	}
	line := LineStringZ(points)
	return NewRingFromLineZ(line)
}

func NewRingFromLineZ(line LineStringZ) *LinearRingZ {
	firstPoint := line[0]
	endPoint := line[len(line)-1]
	if firstPoint.Equal(endPoint) {
		LinearRingZ := LinearRingZ(line)
		return &LinearRingZ
	} else {
		line = append(line, firstPoint)
		LinearRingZ := LinearRingZ(line)
		return &LinearRingZ
	}
}

func (ring LinearRingZ) GetPointCount() int {
	return len(ring)
}

func (ring LinearRingZ) GetPointSet() []PointZ {
	return ring
}

func (ring LinearRingZ) ToLineString() LineStringZ {
	return LineStringZ(ring)
}
