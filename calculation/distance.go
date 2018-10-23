package calculation

import (
	"math"

	"github.com/sadnessly/go-geo/element"
)

//Euclidean distance
func PointDistance(p1 element.Point, p2 element.Point) float64 {
	return math.Sqrt((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))
}

// //计算点到直线的距离 向量的方法，先求三角形的面积，再用面积除以底边长
func PointToLineDistance(point, p1, p2 element.Point) float64 {
	if p1.Equal(p2) {
		return PointDistance(p1, point)
	}
	area := polyArea(*element.NewPolygon(*element.NewLinearRing(*element.NewLine(p1, p2, point))))
	dis := PointDistance(p1, p2)
	return 2 * area / dis
}

// // DistanceFrom returns the distance from the boundary of the geometry in
// // the units of the geometry.
// func DistanceFrom(g orb.Geometry, p orb.Point) float64 {
// 	d, _ := DistanceFromWithIndex(g, p)
// 	return d
// }

// // DistanceFromWithIndex returns the minimum euclidean distance
// // from the boundary of the geometry plus the index of the sub-geometry
// // that was the match.
// func DistanceFromWithIndex(g orb.Geometry, p orb.Point) (float64, int) {
// 	if g == nil {
// 		return math.Inf(1), -1
// 	}

// 	switch g := g.(type) {
// 	case orb.Point:
// 		return Distance(g, p), 0
// 	case orb.MultiPoint:
// 		return multiPointDistanceFrom(g, p)
// 	case orb.LineString:
// 		return lineStringDistanceFrom(g, p)
// 	case orb.MultiLineString:
// 		dist := math.Inf(1)
// 		index := -1
// 		for i, ls := range g {
// 			if d, _ := lineStringDistanceFrom(ls, p); d < dist {
// 				dist = d
// 				index = i
// 			}
// 		}

// 		return dist, index
// 	case orb.Ring:
// 		return lineStringDistanceFrom(orb.LineString(g), p)
// 	case orb.Polygon:
// 		return polygonDistanceFrom(g, p)
// 	case orb.MultiPolygon:
// 		dist := math.Inf(1)
// 		index := -1
// 		for i, poly := range g {
// 			if d, _ := polygonDistanceFrom(poly, p); d < dist {
// 				dist = d
// 				index = i
// 			}
// 		}

// 		return dist, index
// 	case orb.Collection:
// 		dist := math.Inf(1)
// 		index := -1
// 		for i, ge := range g {
// 			if d, _ := DistanceFromWithIndex(ge, p); d < dist {
// 				dist = d
// 				index = i
// 			}
// 		}

// 		return dist, index
// 	case orb.Bound:
// 		return DistanceFromWithIndex(g.ToRing(), p)
// 	}

// 	panic(fmt.Sprintf("geometry type not supported: %T", g))
// }

// func multiPointDistanceFrom(mp element.MultiPoint, p element.Point) (float64, int) {
// 	dist := math.Inf(1)
// 	index := -1

// 	for i := range mp {
// 		if d := DistanceSquared(mp[i], p); d < dist {
// 			dist = d
// 			index = i
// 		}
// 	}

// 	return math.Sqrt(dist), index
// }

func lineStringDistanceFrom(ls element.LineString, p element.Point) (float64, int) {
	dist := math.Inf(1)
	index := -1

	for i := 0; i < len(ls)-1; i++ {
		if d := segmentDistanceFromSquared(ls[i], ls[i+1], p); d < dist {
			dist = d
			index = i
		}
	}

	return math.Sqrt(dist), index
}

func polygonDistanceFrom(p element.Polygon, point element.Point) (float64, int) {
	if len(p) == 0 {
		return math.Inf(1), -1
	}

	dist, index := lineStringDistanceFrom(element.LineString(p[0]), point)
	for i := 1; i < len(p); i++ {
		d, i := lineStringDistanceFrom(element.LineString(p[i]), point)
		if d < dist {
			dist = d
			index = i
		}
	}

	return dist, index
}

func segmentDistanceFromSquared(p1, p2, point element.Point) float64 {
	x := p1.X
	y := p1.Y
	dx := p2.X - x
	dy := p2.Y - y

	if dx != 0 || dy != 0 {
		t := ((point.X-x)*dx + (point.Y-y)*dy) / (dx*dx + dy*dy)

		if t > 1 {
			x = p2.X
			y = p2.Y
		} else if t > 0 {
			x += dx * t
			y += dy * t
		}
	}

	dx = point.X - x
	dy = point.X - y

	return dx*dx + dy*dy
}
