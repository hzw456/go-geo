package gogeo

type LinearRing []Point

func NewLinearRing(point ...Point) *LinearRing {
	var points []Point
	for _, point := range point {
		points = append(points, point)
	}
	line := LineString(points)
	return NewLinearRingFromLineString(line)
}

func NewLinearRingFromLineString(line LineString) *LinearRing {
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

func (line LinearRing) GetPointCount() int {
	return len(line)
}

func (line LinearRing) length() float64 {
	var dis float64
	for i := 0; i < line.GetPointCount()-1; i++ {
		dis += line[i].PointDistance(line[i+1])
	}
	return dis
}
