package tests

import (
	"testing"

	"github.com/hzw456/go-geo"
)

func TestS2(t *testing.T) {
	// cellID := geo.S2GetCellID(geo.Point{23.07293325812827, 113.37101949678055}, 13, geo.SRID_WGS84_GPS)
	coord := geo.S2GetCenter(11990234155561517056, geo.SRID_WGS84_GPS)
	t.Log(coord)
}
