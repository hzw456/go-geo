package geo

type Geometry interface {
	Type() GeometryType
	// BoundingBox() Box
}

//模拟enum类型，对geometry进行枚举
var (
	_ Geometry = Point{}
	_ Geometry = MultiPoint{}
	_ Geometry = LineString{}
	_ Geometry = MultiLineString{}
	_ Geometry = Polygon{}
	_ Geometry = MultiPolygon{}
	// _ Geometry = Collection{}

	//3d geometry
	_ Geometry = PointZ{}
	_ Geometry = MultiPointZ{}
	_ Geometry = LineStringZ{}
	_ Geometry = MultiLineStringZ{}
	_ Geometry = PolygonZ{}
	_ Geometry = MultiPolygonZ{}
)
