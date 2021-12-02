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
	geom, _ := wkt.Decode("POLYGON((23.167959059560499 113.444875331122134,23.167870240520788 113.444956239154266,23.167824392635907 113.444899423736146,23.167918965135584 113.444820493455495,23.167959059560499 113.444875331122134))")
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
