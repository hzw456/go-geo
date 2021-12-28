package tests

import (
	"fmt"
	"testing"

	"github.com/hzw456/go-geo"
	"github.com/hzw456/go-geo/wkt"
)

func TestDistance(t *testing.T) {
	newPoint := geo.NewPoint(153.101112401, 27.797998206)
	newPoint1 := geo.NewPoint(200, 200)
	//newPoint.SetX(500)
	t.Log(geo.PointDistance(*newPoint, *newPoint1, geo.SRID_WGS84_UTM_ZONE_50N))
	//t.Log(geo.Centroid(*newPoint, *newPoint1))
}

func TestLength(t *testing.T) {
	newPoint := geo.NewPoint(153.101112401, 27.797998206)
	newPoint1 := geo.NewPoint(200, 200)
	line := geo.NewLineString(*newPoint, *newPoint1)
	t.Log(line.Length())
}

func TestArea(t *testing.T) {
	newPoint1 := geo.NewPoint(100, 100)
	newPoint2 := geo.NewPoint(200, 100)
	newPoint3 := geo.NewPoint(200, 200)
	newPoint4 := geo.NewPoint(100, 200)
	lr := geo.NewLinearRing(*newPoint1, *newPoint2, *newPoint3, *newPoint4)
	poly := *geo.NewPolygon(*lr)
	t.Log(geo.GetArea(poly))
}

func TestSimplify(t *testing.T) {
	newPoint1 := geo.NewPoint(0, 0)
	newPoint2 := geo.NewPoint(0.5, 0.5)
	newPoint3 := geo.NewPoint(1, 1)
	newPoint4 := geo.NewPoint(2, 2)
	newPoint5 := geo.NewPoint(3, 3)
	line := *geo.NewLineString(*newPoint1, *newPoint2, *newPoint3, *newPoint4, *newPoint5)
	fmt.Println(geo.DouglasPeuckerSimplify(line, 1, geo.SRID_WGS84_UTM_ZONE_44N))
}

func TestPointToLineDis(t *testing.T) {
	newPoint1 := *geo.NewPoint(0, 0)
	newPoint2 := *geo.NewPoint(1, 2)
	newPoint3 := *geo.NewPoint(2, 5)
	fmt.Println(geo.PointToLineDistance(newPoint2, newPoint1, newPoint3, geo.SRID_WGS84_UTM_ZONE_44N))
}

func TestWkt(t *testing.T) {
	newPoint1 := *geo.NewPoint(0, 0)
	newPoint2 := *geo.NewPoint(1, 2)
	newPoint3 := *geo.NewPoint(2, 2)
	newPoint4 := *geo.NewPoint(2, 5)
	line1 := *geo.NewLineString(newPoint1, newPoint2, newPoint3)
	line2 := *geo.NewLineString(newPoint2, newPoint3, newPoint4)
	poly1 := *geo.NewPolygon(*geo.NewRingFromLine(line1))
	poly2 := *geo.NewPolygon(*geo.NewRingFromLine(line2))
	fmt.Println(wkt.Encode(*geo.NewMultiPolygon(poly1, poly2)))
}

func TestEnvolope(t *testing.T) {
	newPoint1 := *geo.NewPoint(0, 0)
	newPoint2 := *geo.NewPoint(1, 2)
	newPoint3 := *geo.NewPoint(2, 2)
	//newPoint4 := *geo.NewPoint(2, 5)
	line1 := *geo.NewLineString(newPoint1, newPoint2, newPoint3)
	//line2 := *geo.NewLine(newPoint2, newPoint3, newPoint4)
	poly1 := *geo.NewPolygon(*geo.NewRingFromLine(line1))
	//poly2 := *geo.NewPolygon(*geo.NewLinearRing(line2))

	fmt.Println(wkt.Encode(geo.BoxToGeo(*geo.BoundingBox(poly1))))
}

func TestPointinPoly(t *testing.T) {
	newPoint1 := *geo.NewPoint(0, 0)
	newPoint2 := *geo.NewPoint(1, 2)
	newPoint3 := *geo.NewPoint(2, 2)
	newPoint4 := *geo.NewPoint(2, 5)
	newPoint5 := *geo.NewPoint(0.5, 0.7)
	newPoint6 := *geo.NewPoint(0.5, 0.5)
	line1 := *geo.NewLineString(newPoint1, newPoint2, newPoint3)
	poly1 := *geo.NewPolygon(*geo.NewRingFromLine(line1))
	if geo.IsPointInPolygon(newPoint4, poly1) == geo.RELA_CONTAIN {
		t.Error("failed, the point is not in poly")
	} else {
		t.Log("success")
	}

	if geo.IsPointInPolygon(newPoint5, poly1) == geo.RELA_DISJOINT {
		t.Error("failed, the point is in poly")
	} else {
		t.Log("success")
	}

	if geo.IsPointInPolygon(newPoint6, poly1) == geo.RELA_CONTAIN {
		t.Error("failed, the point is not in poly")
	} else {
		t.Log("success")
	}
}
