package element

type LinearRing []Point

func NewLinearRing(line LineString) *LinearRing {
	firstPoint, err := line.GetFirstPoint()
	endPoint, err := line.GetEndPoint()
	if err != nil {
		return nil
	}
	if firstPoint.Equal(endPoint) {
		linearRing := LinearRing(line)
		return &linearRing
	} else {
		line = append(line, firstPoint)
		linearRing := LinearRing(line)
		return &linearRing
	}
}

func (line *LinearRing) GetPointCount() int {
	if line == nil {
		return 0
	}
	return len(*line) - 1
}

func (line *LinearRing) length() float64 {
	var dis float64
	for i := 0; i < line.GetPointCount()-1; i++ {
		dis += (*line)[i].PointDistance((*line)[i+1])
	}
	return dis
}
