package tests

import (
	"testing"

	"github.com/hzw456/go-geo"
	"github.com/hzw456/go-geo/wkt"
)

func TestGeohashGetCellID(t *testing.T) {
	// cellID := geo.GeohashGetCellID(geo.Point{23.07293325812827, 113.37101949678055}, 13, geo.SRID_WGS84_GPS)
	coord := geo.GeohashEncode(geo.Point{113.444956239154266, 23.167870240520788}, 12, geo.SRID_WGS84_GPS)
	t.Log(coord)
}

func TestGeohashGetCenter(t *testing.T) {
	// cellID := geo.GeohashGetCellID(geo.Point{23.07293325812827, 113.37101949678055}, 13, geo.SRID_WGS84_GPS)
	coord := geo.GeohashGetCenter("ws0ev984eh43", geo.SRID_WGS84_GPS)
	t.Log(coord)
}

func TestGeohashGetAllNeighbors(t *testing.T) {
	cellID := geo.GeohashGetAllNeighbors("ws0ev984eh43", geo.SRID_WGS84_GPS)
	t.Log(cellID)
}

func TestGeohashboundry(t *testing.T) {
	// cellID := geo.GeohashGetCellID(geo.Point{23.07293325812827, 113.37101949678055}, 13, geo.SRID_WGS84_GPS)
	box := geo.GeohashGetBoundary("ws0ev984eh43", geo.SRID_WGS84_GPS)
	t.Log(box)
	poly := geo.BoxToGeo(box)
	t.Log(wkt.Encode(poly))
}
