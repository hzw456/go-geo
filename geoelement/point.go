package geoelement

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func NewPointStr(point string) {

}

func (p *Point) SetX(x float64) {
	p.X = x
}

func (p *Point) SetY(y float64) {
	p.Y = y
}

func (p Point) Equals(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}
