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

type GeoStringType int16

const (
	STR_GEOJSON GeoStringType = iota
	STR_WKT
	STR_POIJSON
)

type GeometryType string

//定义
const (
	GEOMETRY_POINT            GeometryType = "Point"
	GEOMETRY_MULTIPOINT       GeometryType = "MultiPoint"
	GEOMETRY_LINESTRING       GeometryType = "LineString"
	GEOMETRY_MULTILINESTRING  GeometryType = "MultiLineString"
	GEOMETRY_POLYGON          GeometryType = "Polygon"
	GEOMETRY_MULTIPOLYGON     GeometryType = "MultiPolygon"
	GEOMETRY_COLLECTION       GeometryType = "GeometryCollection"
	GEOMETRY_POINTZ           GeometryType = "PointZ"
	GEOMETRY_MULTIPOINTZ      GeometryType = "MultiPointZ"
	GEOMETRY_LINESTRINGZ      GeometryType = "LineStringZ"
	GEOMETRY_MULTILINESTRINGZ GeometryType = "MultiLineStringZ"
	GEOMETRY_POLYGONZ         GeometryType = "PolygonZ"
	GEOMETRY_MULTIPOLYGONZ    GeometryType = "MultiPolygonZ"
	GEOMETRY_COLLECTIONZ      GeometryType = "GeometryCollectionZ"
)

type SRID int

const (
	SRID_WGS84_GPS             SRID = 4326 //wgs84
	SRID_WGS84_PSEUDO_MERCATOR SRID = 3857 //wgs84,Pseudo-Mercator
	SRID_WGS84_UTM_ZONE_44N    SRID = 32644
	SRID_WGS84_UTM_ZONE_45N    SRID = 32645 //utm
	SRID_WGS84_UTM_ZONE_46N    SRID = 32646 //utm
	SRID_WGS84_UTM_ZONE_47N    SRID = 32647 //utm
	SRID_WGS84_UTM_ZONE_48N    SRID = 32648 //utm
	SRID_WGS84_UTM_ZONE_49N    SRID = 32649 //utm
	SRID_WGS84_UTM_ZONE_50N    SRID = 32650 //utm
	SRID_WGS84_UTM_ZONE_51N    SRID = 32651 //utm
	SRID_WGS84_UTM_ZONE_52N    SRID = 32652 //utm
	SRID_WGS84_UTM_ZONE_53N    SRID = 32653 //utm

)
