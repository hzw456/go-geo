package geo

type LineStringZ []PointZ

func NewLineStringZ(point ...PointZ) *LineStringZ {
	var points []PointZ
	for _, point := range point {
		points = append(points, point)
	}
	line := LineStringZ(points)
	return &line
}

func (line LineStringZ) GetPointCount() int {
	return len(line)
}

//取首点
func (line LineStringZ) GetFirstPoint() PointZ {
	if len(line) < 2 {
		return PointZ{0, 0, 0}
	}
	return line[0]
}

//取尾点
func (line LineStringZ) GetEndPoint() PointZ {
	if len(line) < 2 {
		return PointZ{0, 0, 0}
	}
	return line[line.GetPointCount()-1]
}

func (line *LineStringZ) Append(point PointZ) error {
	*line = append(*line, point)
	return nil
}

//判断两条线是不是相同
func (line1 LineStringZ) Equal(line2 LineStringZ) bool {
	for i := 0; i < line1.GetPointCount(); i++ {
		if !line1[i].Equal(line2[i]) {
			return false
		}
	}
	return true
}

func (line *LineStringZ) Reverse() {
	count := line.GetPointCount()
	mid := count / 2
	for i := 0; i < mid; i++ {
		(*line)[i], (*line)[line.GetPointCount()-1-i] = (*line)[line.GetPointCount()-1-i], (*line)[i]
	}
}

func (line LineStringZ) ToRing() LinearRingZ {
	return *NewRingFromLineZ(line)
}

func (line LineStringZ) GetPointSet() []PointZ {
	return line
}
