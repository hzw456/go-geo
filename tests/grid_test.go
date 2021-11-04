package tests

import (
	"testing"

	"github.com/hzw456/go-geo"
	"github.com/hzw456/go-geo/wkt"
)

func TestS2GetCellID(t *testing.T) {
	// cellID := geo.S2GetCellID(geo.Point{23.07293325812827, 113.37101949678055}, 13, geo.SRID_WGS84_GPS)
	coord := geo.S2GetCellID(geo.Point{23.167870240520788, 113.444956239154266}, 13, geo.SRID_WGS84_GPS)
	t.Log(coord)
}

func TestS2GetCenter(t *testing.T) {
	// cellID := geo.S2GetCellID(geo.Point{23.07293325812827, 113.37101949678055}, 13, geo.SRID_WGS84_GPS)
	coord := geo.S2GetCenter(3747843558330073088, geo.SRID_WGS84_GPS)
	t.Log(coord)
}

func TestS2GetAllNeighbors(t *testing.T) {
	cellID := geo.S2GetAllNeighbors(11990234155561517056, geo.SRID_WGS84_GPS)
	t.Log(cellID)
}

func TestS2RegionCoverer(t *testing.T) {
	geom, _ := wkt.Decode("POLYGON((23.171206994915124 113.442910973440775,23.171141623923660 113.442952310979493,23.171207269583153 113.442911110774787,23.171230616365818 113.442954096321699,23.171165657376402 113.442995433860418))")
	poly := geom.(geo.Polygon)
	t.Log(wkt.Encode(poly))
	cellID := geo.RegionCoverer(poly, 22, geo.SRID_WGS84_GPS)
	t.Log(cellID)
}

func TestS2boundry(t *testing.T) {
	// cellID := geo.S2GetCellID(geo.Point{23.07293325812827, 113.37101949678055}, 13, geo.SRID_WGS84_GPS)
	coord := geo.S2GetBoundary(3747843557077024768, geo.SRID_WGS84_GPS)
	t.Log(coord)
	poly := geo.NewPolygonFromPois(coord...)
	t.Log(wkt.Encode(*poly))
}
