package gogeo

func segmentIntersect(p0, p1, p2, p3 Point) bool {
	s1X := p1.X - p0.X
	s1Y := p1.Y - p0.Y
	s2X := p3.X - p2.X
	s2Y := p3.Y - p2.Y

	s := (-s1Y*(p0.X-p2.X) + s1X*(p0.Y-p2.Y)) / (-s2X*s1Y + s1X*s2Y)
	t := (s2X*(p0.Y-p2.Y) - s2Y*(p0.X-p2.X)) / (-s2X*s1Y + s1X*s2Y)

	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
		return true
	}
	return false // No collision
}
