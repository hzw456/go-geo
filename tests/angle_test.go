package tests

import (
	"testing"

	"github.com/hzw456/go-geo"
)

func TestGetAzimuth(t *testing.T) {
	newPoint1 := *geo.NewPoint(22.956205800445613, 112.95067837273719)
	newPoint2 := *geo.NewPoint(23.045253006300054, 113.0495284775218)
	t.Log(geo.GetAzimuth(newPoint1, newPoint2, geo.SRID_WGS84_GPS))
}
