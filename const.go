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
	GEO_EQUAL     = 3
	GEO_TOUCH     = 4
	GEO_COVER     = 5
	GEO_INTERSECT = 6
)
