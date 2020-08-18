package geo

const (
	COORDPRESION  = 0.000001
	INF           = float64(1 << 31)
	EARTHRADIUSMI = 6371000
	EARTHRADIUSKM = 6371
)

type GeometryRealation int16

//定义空间关系类型
const (
	RELA_UNKNOWN GeometryRealation = iota
	RELA_DISJOINT
	RELA_CONTAIN
	RELA_EQUAL
	RELA_TOUCH
	RELA_COVER
	RELA_INTERSECT
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

type SRID int

const (
	SRID_WGS84_GPS             SRID = 4326 //wgs84
	SRID_WGS84_PSEUDO_MERCATOR SRID = 3824 //wgs84,Pseudo-Mercator

)

type GeoJSONType string

const (
	PointType              GeoJSONType = "Point"
	MultiPointType         GeoJSONType = "MultiPoint"
	LineStringType         GeoJSONType = "LineString"
	MultiLineStringType    GeoJSONType = "MultiLineString"
	PolygonType            GeoJSONType = "Polygon"
	MultiPolygonType       GeoJSONType = "MultiPolygon"
	GeometryCollectionType GeoJSONType = "GeometryCollection"
	FeatureType            GeoJSONType = "Feature"
	FeatureCollectionType  GeoJSONType = "FeatureCollection"
)
