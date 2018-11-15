package tests

import (
	"testing"

	gogeo "github.com/sadnessly/go-geo"
)

func TestHexCentriod(t *testing.T) {
	newPoint1 := *gogeo.NewPoint(0, 0)
	newPoint2 := *gogeo.NewPoint(1, 1)
	newPoint3 := *gogeo.NewPoint(2, 0)
	line1 := *gogeo.NewLine(newPoint1, newPoint2, newPoint3)
	poly1 := *gogeo.NewPolygon(*gogeo.NewLinearRing(line1))
	t.Log(gogeo.GeoToWkt(gogeo.Centroid(poly1)))
}
