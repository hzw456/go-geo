package simplify

type Simplifier interface {
	//Equal(geo Geometry) bool
}

var (
	_ Simplifier = DouglasPeuckerSimplifier{}
)
