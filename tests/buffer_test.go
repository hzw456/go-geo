package tests

import (
	"testing"

	gogeo "github.com/sadnessly/go-geo"
)

func TestPointBuffer(t *testing.T) {
	newPoint1 := gogeo.NewPoint(200, 200)
	gogeo.Buffer(*newPoint1, 10)
	t.Log(gogeo.GeoToWkt(gogeo.Buffer(*newPoint1, 10)))
	newLine := gogeo.NewLine(gogeo.Point{0, 0}, gogeo.Point{5, 2}, gogeo.Point{7, 2}, gogeo.Point{10, 12})
	t.Log(gogeo.GeoToWkt(gogeo.Buffer(*newLine, 5)))
}

func TestPolyBuffer(t *testing.T) {
	newPoint1 := *gogeo.NewPoint(0, 0)
	newPoint2 := *gogeo.NewPoint(1, 1)
	newPoint3 := *gogeo.NewPoint(2, 0)
	poly1 := *gogeo.NewPolygonFromPois(newPoint1, newPoint2, newPoint3)
	t.Log(gogeo.GeoToWkt(poly1.Buffer(5)))
}
