package gogeo

const (
	COORDPRESION = 0.00000001
	INF          = float64(1 << 31)
)

type GeometryRealation int16

//定义空间关系类型
const (
	GEO_UNKNOWN GeometryRealation = iota
	GEO_DISJOINT
	GEO_CONTAIN
	GEO_EQUAL
	GEO_TOUCH
	GEO_COVER
	GEO_INTERSECT
)

type GeometryType int16

//定义
const (
	ELEM_POINT GeometryType = iota
	ELEM_MULTIPOINT
	ELEM_LINESTRING
	ELEM_MULTILINESTRING
	ELEM_LINEARRING
	ELEM_POLYGON
	ELEM_MULTIPOLYGON
	ELEM_COLLECTION
	ELEM_UNKNOWN
)

type GeoStringType int16

const (
	STR_GEOJSON GeoStringType = iota
	STR_WKT
	STR_POIJSON
)
