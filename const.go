package gogeo

const (
	COORDPRESION = 0.00000001
	INF          = float64(1 << 31)
)

//定义空间关系类型
const (
	GEO_UNKNOWN   = 0
	GEO_DISJOINT  = 1
	GEO_CONTAIN   = 2
	GEO_CONTAINBY = 3
	GEO_EQUAL     = 4
	GEO_TOUCH     = 5
	GEO_COVER     = 6
	GEO_COVERBY   = 7
	GEO_INTERSECT = 8
)
