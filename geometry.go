package geo

type Geometry interface {
	ToWkt() ([]byte, error)
	BoundingBox() Box
	Type() string
	ToGeojson() ([]byte, error)
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
)

type GeometryZ interface {
	ToGeojson() ([]byte, error)
}

//模拟enum类型，对geometry进行枚举
var (
	_ GeometryZ = PointZ{}
	_ GeometryZ = MultiPointZ{}
	_ GeometryZ = LineStringZ{}
	_ GeometryZ = MultiLineStringZ{}
	_ GeometryZ = PolygonZ{}
	_ GeometryZ = MultiPolygonZ{}
	// _ GeometryZ = CollectionZ{}
)

type GeometryM interface {
}

type GeometryZM interface {
}
