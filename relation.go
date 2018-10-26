package gogeo

func euqal(x float64, y float64) bool {
	v := x - y
	const delta float64 = 1e-6
	if v < delta && v > -delta {
		return true
	}
	return false

}

func little(x float64, y float64) bool {
	if euqal(x, y) {
		return false
	}
	return x < y
}

func little_equal(x float64, y float64) bool {
	if euqal(x, y) {
		return true
	}
	return x < y
}

func IsPointInPolygon(point Point, poly Polygon) bool {
	isIn := false

	lr := poly.GetExteriorRing()
	pointCount := lr.GetPointCount()
	for i := 0; i < pointCount; i++ {
		j := i - 1
		if i == 0 {
			j = pointCount - 1
		}
		vi := lr[i]
		vj := lr[j]

		xmin := vi.X
		xmax := vj.X
		if xmin > xmax {
			t := xmin
			xmin = xmax
			xmax = t
		}
		ymin := vi.Y
		ymax := vj.Y
		if ymin > ymax {
			t := ymin
			ymin = ymax
			ymax = t
		}
		// i//j//aixs_x
		if euqal(vj.Y, vi.Y) {
			if euqal(point.Y, vi.Y) && little_equal(xmin, point.X) && little_equal(point.X, xmax) {
				return true
			}
			continue
		}

		xt := (vj.X-vi.X)*(point.Y-vi.Y)/(vj.Y-vi.Y) + vi.X
		if euqal(xt, point.X) && little_equal(ymin, point.Y) && little_equal(point.Y, ymax) {
			// on edge [vj,vi]
			return true
		}
		if little(point.X, xt) && little_equal(ymin, point.Y) && little(point.Y, ymax) {
			isIn = !isIn
		}

	}
	return isIn
}
