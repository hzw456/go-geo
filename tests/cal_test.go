package tests

import (
	"math"
	"testing"

	"github.com/hzw456/go-geo"
	"github.com/hzw456/go-geo/wkt"
)

func TestCentriod(t *testing.T) {
	newPoint1 := *geo.NewPoint(0, 0)
	newPoint2 := *geo.NewPoint(1, 1)
	newPoint3 := *geo.NewPoint(2, 0)
	line1 := *geo.NewLineString(newPoint1, newPoint2, newPoint3)
	poly1 := *geo.NewPolygon(*geo.NewRingFromLine(line1))
	t.Log(wkt.Encode(geo.Centroid(poly1)))
}

func TestRotate(t *testing.T) {
	newPoint1 := *geo.NewPoint(0, 0)
	newPoint2 := *geo.NewPoint(1, 1)
	newPoint3 := *geo.NewPoint(2, 0)
	pt1 := geo.RotateCCW(newPoint2, newPoint1, 45.0/180.0*math.Pi)
	t.Log(wkt.Encode(pt1))
	poly1 := *geo.NewPolygonFromPois(newPoint1, newPoint2, newPoint3)
	poly2 := geo.RotateCCW(poly1, newPoint1, 45.0/180.0*math.Pi)
	t.Log(wkt.Encode(poly2))
}

func TestHausdorff(t *testing.T) {
	newPoint1 := *geo.NewPoint(0, 0)
	newPoint2 := *geo.NewPoint(1, 1)
	newPoint3 := *geo.NewPoint(2, 0)
	newPoint4 := *geo.NewPoint(3, 0)
	line1 := *geo.NewLineString(newPoint1, newPoint2, newPoint3)
	line2 := *geo.NewLineString(newPoint1, newPoint2, newPoint4)
	dis := geo.HausdorffDistance(line1, line2, geo.SRID_WGS84_PSEUDO_MERCATOR)
	t.Log(dis)
}

func TestGetArea(t *testing.T) {
	newPoint1 := *geo.NewPoint(0, 0)
	newPoint2 := *geo.NewPoint(1, 1)
	newPoint3 := *geo.NewPoint(2, 0)
	line1 := *geo.NewLineString(newPoint1, newPoint2, newPoint3)
	poly1 := *geo.NewPolygon(*geo.NewRingFromLine(line1))
	t.Log(geo.GetArea(poly1))
}
