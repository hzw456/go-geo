package geos

type Point struct {
	X, Y float64
}

type CoordZ struct {
	X, Y, Z float64
}

func (p *Point) Equals(p2 *Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (p *CoordZ) Equals(p2 *CoordZ) bool {
	return p.X == p2.X && p.Y == p2.Y && p.Z == p2.Z
}
