package geo

type Geometry interface {
	SetSrid(srid uint64)
	// GetSRID() uint64
	ToWkt() string
	BoundingBox() Box
	// Buffer(dis float64) Polygon
}

type Collection []Geometry

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
