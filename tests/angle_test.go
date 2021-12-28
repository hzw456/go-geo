package tests

import (
	"testing"

	"github.com/hzw456/go-geo"
)

func TestGetAzimuth(t *testing.T) {
	newPoint1 := *geo.NewPoint(112.95067837273719, 22.956205800445613)
	newPoint2 := *geo.NewPoint(113.0495284775218, 23.045253006300054)
	t.Log(geo.GetAzimuth(newPoint1, newPoint2, geo.SRID_WGS84_GPS))
}
