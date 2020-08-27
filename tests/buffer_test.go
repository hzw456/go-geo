package tests

import (
	"fmt"
	"testing"

	"github.com/sadnessly/go-geo"
)

// func TestPointBuffer(t *testing.T) {
// 	newPoint1 := geo.NewPoint(200, 200)
// 	geo.Buffer(*newPoint1, 10)
// 	t.Log(geo.GeoToWkt(geo.Buffer(*newPoint1, 10)))
// 	newLine := geo.NewLine(geo.Point{0, 0}, geo.Point{5, 2}, geo.Point{7, 2}, geo.Point{10, 12})
// 	t.Log(geo.GeoToWkt((*newLine, 5)))
// }

func TestPolyBuffer(t *testing.T) {
	newPoint1 := *geo.NewPoint(0, 0)
	newPoint2 := *geo.NewPoint(1, 1)
	newPoint3 := *geo.NewPoint(2, 0)
	poly1 := *geo.NewPolygonFromPois(newPoint1, newPoint2, newPoint3)
	w, err := poly1.Buffer(5).ToWkt()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(w))
}

func TestPolyBuffer2(t *testing.T) {
	newPoint1 := *geo.NewPointZ(0, 0, 0)
	newPoint2 := *geo.NewPointZ(1, 1, 0)
	newPoint3 := *geo.NewPointZ(2, 0, 0)
	poly1 := *geo.NewPolygonZFromPois(newPoint1, newPoint2, newPoint3)
	w, err := poly1.ToGeojson()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(w))
}
