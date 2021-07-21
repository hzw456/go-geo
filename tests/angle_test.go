package tests

import (
	"testing"

	"github.com/hzw456/go-geo"
)

func TestGetAzimuth(t *testing.T) {
	newPoint1 := *geo.NewPoint(23.308908764595763, 101.5422887717471)
	newPoint2 := *geo.NewPoint(23.308749916024873, 101.55205930972208)
	t.Log(geo.GetAzimuth(newPoint1, newPoint2, geo.SRID_WGS84_GPS))
}
