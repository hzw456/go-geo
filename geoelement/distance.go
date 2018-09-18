package geoelement

import "math"

//Euclidean distance
func pointDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(p1.X*p1.X + p1.Y*p1.Y)
}
