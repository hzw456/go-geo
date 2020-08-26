package geo

type Geometry interface {
	SetSrid(srid uint64)
	// GetSRID() uint64
	ToWkt() ([]byte, error)
	BoundingBox() Box
	Type() string
	ToGeojson() ([]byte, error)
	// Buffer(dis float64) Polygon
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
	SetSrid(srid uint64)
	// GetSRID() uint64
	ToWkt() ([]byte, error)
	BoundingBox() Box
	Type() string
	ToGeojson() ([]byte, error)
	// Buffer(dis float64) Polygon
}

// //模拟enum类型，对geometry进行枚举
// var (
// 	_ GeometryZ = PointZ{}
// 	_ GeometryZ = MultiPointZ{}
// 	_ GeometryZ = LineStringZ{}
// 	_ GeometryZ = MultiLineStringZ{}
// 	_ GeometryZ = PolygonZ{}
// 	_ GeometryZ = MultiPolygonZ{}
// 	// _ GeometryZ = CollectionZ{}
// )

type GeometryM interface {
	SetSrid(srid uint64)
	// GetSRID() uint64
	ToWkt() ([]byte, error)
	BoundingBox() Box
	Type() string
	ToGeojson() ([]byte, error)
	// Buffer(dis float64) Polygon
}

type GeometryZM interface {
	SetSrid(srid uint64)
	// GetSRID() uint64
	ToWkt() ([]byte, error)
	BoundingBox() Box
	Type() string
	ToGeojson() ([]byte, error)
	// Buffer(dis float64) Polygon
}
