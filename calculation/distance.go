package calculation

import (
	"math"

	"git.xiaojukeji.com/haozhiwei/go-geo/element"
)

//Euclidean distance
func PointDistance(p1 element.Point, p2 element.Point) float64 {
	return math.Sqrt(p1.X*p1.X + p1.Y*p1.Y)
}

// //point to line distance
// func PointToLineDistance() float64 {

// }
