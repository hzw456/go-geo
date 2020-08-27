package geo

func (p Point) Type() string {
	return "Point"
}

func (p PointZ) Type() string {
	return "Point"
}

func (mp MultiPoint) Type() string {
	return "MultiPoint"
}

func (mp MultiPointZ) Type() string {
	return "MultiPoint"
}

func (l LineString) Type() string {
	return "LineString"
}

func (l LineStringZ) Type() string {
	return "LineString"
}

func (ml MultiLineString) Type() string {
	return "MultiLineString"
}

func (ml MultiLineStringZ) Type() string {
	return "MultiLineString"
}

func (mp MultiPolygon) Type() string {
	return "MultiPolygon"
}

func (mp MultiPolygonZ) Type() string {
	return "MultiPolygon"
}

func (p Polygon) Type() string {
	return "Polygon"
}

func (p PolygonZ) Type() string {
	return "Polygon"
}

func (c Collection) Type() string {
	return "GeometryCollection"
}
