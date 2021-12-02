package geo

import (
	"math"
	"sort"
)

type hullPts []Point

//Implement sort interface
func (p hullPts) Len() int {
	return len(p)
}

func (p hullPts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p hullPts) Less(i, j int) bool {
	area := Area2(p[0], p[i], p[j])

	if area == 0 {
		x := math.Abs(p[i].X-p[0].X) - math.Abs(p[j].X-p[0].X)
		y := math.Abs(p[i].Y-p[0].Y) - math.Abs(p[j].Y-p[0].Y)

		if x < 0 || y < 0 {
			return true
		} else if x > 0 || y > 0 {
			return false
		} else {
			return false
		}
	}

	return area > 0
}

func (p hullPts) FindLowestPoint() {
	m := 0
	for i := 1; i < len(p); i++ {
		//If lowest points are on the same line, take the rightmost point
		if (p[i].Y < p[m].Y) || ((p[i].Y == p[m].Y) && p[i].X > p[m].X) {
			m = i
		}
	}
	p[0], p[m] = p[m], p[0]
}

func ConvexHull(pts ...Point) Polygon {
	if len(pts) < 3 {
		return nil
	}
	points := hullPts(pts)
	var stack hullPts
	points.FindLowestPoint()
	sort.Sort(&points)
	stack = append(stack, points[0])
	stack = append(stack, points[1])
	i := 2
	for i < len(points) {
		pi := points[i]

		p1 := stack[len(stack)-2]
		p2 := stack[len(stack)-1]

		if isLeft(p1, p2, pi) {
			stack = append(stack, pi)
			i++
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return *NewPolygonFromPois(stack...)
}

func isLeft(p0, p1, p2 Point) bool {
	return Area2(p0, p1, p2) >= 0
}

func Area2(a, b, c Point) float64 {
	return (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)
}
