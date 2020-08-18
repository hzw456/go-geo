package geo

type Geometry interface {
	// SetSRID(srid int)
	// ToWkt() string
	// ToGeojson() string
}

type Collection []Geometry

//模拟enum类型，对geometry进行枚举
var (
	_ Geometry = Point{}
	_ Geometry = MultiPoint{}
	_ Geometry = LineString{}
	_ Geometry = MultiLineString{}
	_ Geometry = LinearRing{}
	_ Geometry = Polygon{}
	_ Geometry = MultiPolygon{}
	_ Geometry = Collection{}
)
