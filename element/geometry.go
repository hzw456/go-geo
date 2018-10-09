package element

type geometry interface {
	Buffer(bufferDis float64) Polygon
	Equal() bool
}

type lineElement interface {
	Length() float64
	GetPointCount() int
}

type polyElement interface {
	Area() float64
}
