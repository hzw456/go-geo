package geo

func (p Point) Type() GeometryType {
	return GEOMETRY_POINT
}

func (p PointZ) Type() GeometryType {
	return GEOMETRY_POINTZ
}

func (mp MultiPoint) Type() GeometryType {
	return GEOMETRY_MULTIPOINT
}

func (mp MultiPointZ) Type() GeometryType {
	return GEOMETRY_MULTIPOINTZ
}

func (l LineString) Type() GeometryType {
	return GEOMETRY_LINESTRING
}

func (l LineStringZ) Type() GeometryType {
	return GEOMETRY_LINESTRINGZ
}

func (ml MultiLineString) Type() GeometryType {
	return GEOMETRY_MULTILINESTRING
}

func (ml MultiLineStringZ) Type() GeometryType {
	return GEOMETRY_MULTILINESTRINGZ
}

func (mp MultiPolygon) Type() GeometryType {
	return GEOMETRY_MULTIPOLYGON
}

func (mp MultiPolygonZ) Type() GeometryType {
	return GEOMETRY_MULTIPOLYGONZ
}

func (p Polygon) Type() GeometryType {
	return GEOMETRY_POLYGON
}

func (p PolygonZ) Type() GeometryType {
	return GEOMETRY_POLYGONZ
}

func (c Collection) Type() GeometryType {
	return GEOMETRY_COLLECTION
}
