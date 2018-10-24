package box

import "github.com/sadnessly/go-geo/element"

type box struct {
	minX float64
	minY float64
	maxX float64
	maxY float64
}

func boxToGeo(b box) element.Geometry {
	p1 := element.Point{b.minX, b.minY}
	p2 := element.Point{b.minX, b.maxY}
	p3 := element.Point{b.maxX, b.maxY}
	p4 := element.Point{b.maxX, b.minY}

	if p1.Equal(p3) {
		//元素是个点
		return p1
	} else if p1.Equal(p2) {
		//元素是条线 y坐标不同
		return element.LineString{p1, p3}
	} else if p2.Equal(p3) {
		//元素是条线 x坐标不同
		return element.LineString{p1, p2}
	}
	return *element.NewPolygon(element.LinearRing{p1, p2, p3, p4})
}
