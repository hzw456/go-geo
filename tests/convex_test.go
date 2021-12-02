package tests

import (
	"testing"

	"github.com/hzw456/go-geo"
	"github.com/hzw456/go-geo/wkt"
)

func TestMbr(t *testing.T) {
	newPoint1 := geo.NewPoint(0, 0)
	newPoint2 := geo.NewPoint(0.0, 0.5)
	newPoint3 := geo.NewPoint(1, 1)
	newPoint4 := geo.NewPoint(0, 3)
	newPoint5 := geo.NewPoint(0, 2)
	line := geo.NewMultiPoint(*newPoint1, *newPoint2, *newPoint3, *newPoint4, *newPoint5)
	geom := geo.Mbr(*line)
	t.Log(wkt.Encode(geom))
}

func TestConvexHull(t *testing.T) {
	newPoint1 := geo.NewPoint(0, 0)
	newPoint2 := geo.NewPoint(0.0, 0.5)
	newPoint3 := geo.NewPoint(1, 1)
	newPoint4 := geo.NewPoint(0, 3)
	newPoint5 := geo.NewPoint(0, 2)
	geom := geo.ConvexHull(*newPoint1, *newPoint2, *newPoint3, *newPoint4, *newPoint5)
	t.Log(wkt.Encode(geom))
}
