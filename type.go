package geo

func (p Point) Type() string {
	return "Point"
}

func (mp MultiPoint) Type() string {
	return "MultiPoint"
}

func (l LineString) Type() string {
	return "LineString"
}

func (ml MultiLineString) Type() string {
	return "MultiLineString"
}

func (mp MultiPolygon) Type() string {
	return "MultiPolygon"
}

func (p Polygon) Type() string {
	return "Polygon"
}

func (c Collection) Type() string {
	return "GeometryCollection"
}
