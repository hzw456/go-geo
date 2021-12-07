package tests

import (
	"fmt"
	"testing"

	"github.com/hzw456/go-geo"
)

func TestRTree(t *testing.T) {
	newPoint1 := geo.Spatial{"1", nil, *geo.NewPoint(0, 0)}
	newPoint2 := geo.Spatial{"2", nil, *geo.NewPoint(1, 1)}
	newPoint3 := geo.Spatial{"3", nil, *geo.NewPoint(2, 1)}
	newPoint4 := geo.Spatial{"4", nil, *geo.NewPoint(2, 3)}
	newPoint5 := geo.Spatial{"5", nil, *geo.NewPoint(2, 2)}
	rtree := geo.NewTree(newPoint1, newPoint2, newPoint3, newPoint4, newPoint5)
	sps := rtree.SearchIntersect(geo.NewRectBox(geo.Point{0, 0}, 10, 10))
	fmt.Println(sps)
	rtree.DeleteBySpatialID("1")
	sps = rtree.SearchIntersect(geo.NewRectBox(geo.Point{0, 0}, 10, 10))
	fmt.Println(sps)
	fmt.Println(rtree.Size())
}
