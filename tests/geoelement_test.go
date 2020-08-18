package tests

import (
	"fmt"
	"testing"

	gogeo "github.com/sadnessly/go-geo"
)

func TestDistance(t *testing.T) {
	newPoint := gogeo.NewPoint(153.101112401, 27.797998206)
	newPoint1 := gogeo.NewPoint(200, 200)
	//newPoint.SetX(500)
	t.Log(gogeo.PointDistance(*newPoint, *newPoint1))
	//t.Log(gogeo.Centroid(*newPoint, *newPoint1))
}

func TestLength(t *testing.T) {
	newPoint := gogeo.NewPoint(153.101112401, 27.797998206)
	newPoint1 := gogeo.NewPoint(200, 200)
	line := gogeo.NewLine(*newPoint, *newPoint1)
	t.Log(line.Length())
}

func TestArea(t *testing.T) {
	newPoint1 := gogeo.NewPoint(100, 100)
	newPoint2 := gogeo.NewPoint(200, 100)
	newPoint3 := gogeo.NewPoint(200, 200)
	newPoint4 := gogeo.NewPoint(100, 200)
	lr := gogeo.NewLinearRing(*newPoint1, *newPoint2, *newPoint3, *newPoint4)
	poly := *gogeo.NewPolygon(*lr)
	t.Log(gogeo.GetArea(poly))
}

func TestBuffer(t *testing.T) {
	newPoint1 := gogeo.NewPoint(200, 200)
	gogeo.Buffer(*newPoint1, 10)
	t.Log(gogeo.GeoToWkt(gogeo.Buffer(*newPoint1, 10)))
	newLine := gogeo.NewLine(gogeo.Point{0, 0}, gogeo.Point{5, 2}, gogeo.Point{7, 2}, gogeo.Point{10, 12})
	t.Log(gogeo.GeoToWkt(gogeo.Buffer(*newLine, 5)))
}

func TestSimplify(t *testing.T) {
	newPoint1 := gogeo.NewPoint(0, 0)
	newPoint2 := gogeo.NewPoint(0.5, 0.5)
	newPoint3 := gogeo.NewPoint(1, 1)
	newPoint4 := gogeo.NewPoint(2, 2)
	newPoint5 := gogeo.NewPoint(3, 3)
	line := *gogeo.NewLine(*newPoint1, *newPoint2, *newPoint3, *newPoint4, *newPoint5)
	fmt.Println(gogeo.DouglasPeuckerSimplifier{1}.Simplify(line))
}

func TestPointToLineDis(t *testing.T) {
	newPoint1 := *gogeo.NewPoint(0, 0)
	newPoint2 := *gogeo.NewPoint(1, 2)
	newPoint3 := *gogeo.NewPoint(2, 5)
	fmt.Println(gogeo.PointToLineDistance(newPoint2, newPoint1, newPoint3))
}

func TestWkt(t *testing.T) {
	newPoint1 := *gogeo.NewPoint(0, 0)
	newPoint2 := *gogeo.NewPoint(1, 2)
	newPoint3 := *gogeo.NewPoint(2, 2)
	newPoint4 := *gogeo.NewPoint(2, 5)
	fmt.Println(gogeo.GeoToWkt(gogeo.MultiPoint{newPoint1, newPoint2, newPoint3}))
	line1 := *gogeo.NewLine(newPoint1, newPoint2, newPoint3)
	line2 := *gogeo.NewLine(newPoint2, newPoint3, newPoint4)
	poly1 := *gogeo.NewPolygon(*gogeo.NewLinearRingFromLineString(line1))
	poly2 := *gogeo.NewPolygon(*gogeo.NewLinearRingFromLineString(line2))
	fmt.Println(gogeo.GeoToWkt(*gogeo.NewMultiPolygon(poly1, poly2)))
}

func TestEnvolope(t *testing.T) {
	newPoint1 := *gogeo.NewPoint(0, 0)
	newPoint2 := *gogeo.NewPoint(1, 2)
	newPoint3 := *gogeo.NewPoint(2, 2)
	//newPoint4 := *gogeo.NewPoint(2, 5)
	line1 := *gogeo.NewLine(newPoint1, newPoint2, newPoint3)
	//line2 := *gogeo.NewLine(newPoint2, newPoint3, newPoint4)
	poly1 := *gogeo.NewPolygon(*gogeo.NewLinearRingFromLineString(line1))
	//poly2 := *gogeo.NewPolygon(*gogeo.NewLinearRing(line2))

	fmt.Println(gogeo.GeoToWkt(gogeo.BoxToGeo(gogeo.Envelope(poly1))))
}

func TestPointinPoly(t *testing.T) {
	newPoint1 := *gogeo.NewPoint(0, 0)
	newPoint2 := *gogeo.NewPoint(1, 2)
	newPoint3 := *gogeo.NewPoint(2, 2)
	newPoint4 := *gogeo.NewPoint(2, 5)
	newPoint5 := *gogeo.NewPoint(0.5, 0.7)
	newPoint6 := *gogeo.NewPoint(0.5, 0.5)
	line1 := *gogeo.NewLine(newPoint1, newPoint2, newPoint3)
	poly1 := *gogeo.NewPolygon(*gogeo.NewLinearRingFromLineString(line1))
	if gogeo.IsPointInPolygon(newPoint4, poly1) == gogeo.RELA_CONTAIN {
		t.Error("failed, the point is not in poly")
	} else {
		t.Log("success")
	}

	if gogeo.IsPointInPolygon(newPoint5, poly1) == gogeo.RELA_DISJOINT {
		t.Error("failed, the point is in poly")
	} else {
		t.Log("success")
	}

	if gogeo.IsPointInPolygon(newPoint6, poly1) == gogeo.RELA_CONTAIN {
		t.Error("failed, the point is not in poly")
	} else {
		t.Log("success")
	}
}
