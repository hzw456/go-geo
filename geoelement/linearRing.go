package geoelement

type LinearRing []Point

func (points Points) newLine() LineString {
	return (LineString)points
}
