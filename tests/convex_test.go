package tests

import (
	"testing"

	"github.com/hzw456/go-geo"
	"github.com/hzw456/go-geo/wkt"
)

func TestMbr(t *testing.T) {
	newPoint1 := geo.NewPoint(0, 0)
	newPoint2 := geo.NewPoint(1, 1)
	newPoint3 := geo.NewPoint(2, 1)
	newPoint4 := geo.NewPoint(2, 3)
	newPoint5 := geo.NewPoint(2, 2)
	line := geo.NewMultiPoint(*newPoint1, *newPoint2, *newPoint3, *newPoint4, *newPoint5)
	t.Log(wkt.Encode(*line))
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
func TestMbr2(t *testing.T) {
	geom, _ := wkt.Decode("MULTIPOINT(759345.2410822669 2557684.7848768234,759347.8335629855 2557685.4528085375,759350.4089494486 2557686.116336054)")
	t.Log(wkt.Encode(geom))
	geom2 := geo.Mbr(geom)
	t.Log(wkt.Encode(geom2))
}
