package geo

type LinearRing []Point

func NewLinearRing(point ...Point) *LinearRing {
	var points []Point
	for _, point := range point {
		points = append(points, point)
	}
	line := LineString(points)
	return NewRingFromLine(line)
}

func NewRingFromLine(line LineString) *LinearRing {
	err := line.Verify()
	if err != nil {
		return nil
	}
	firstPoint := line[0]
	endPoint := line[len(line)-1]
	if firstPoint.Equal(endPoint) {
		linearRing := LinearRing(line)
		return &linearRing
	} else {
		line = append(line, firstPoint)
		linearRing := LinearRing(line)
		return &linearRing
	}
}

func (ring LinearRing) GetPointCount() int {
	return len(ring)
}

func (ring LinearRing) length() float64 {
	var dis float64
	for i := 0; i < ring.GetPointCount()-1; i++ {
		dis += ring[i].Distance(ring[i+1])
	}
	return dis
}

func (ring LinearRing) GetPointSet() []Point {
	return ring
}

func (ring LinearRing) ToLineString() LineString {
	return LineString(ring)
}
