package geo

import (
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
)

func S2GetCellID(p Point, level int, srid SRID) uint64 {
	if srid == SRID_WGS84_GPS {
		cellID := s2.CellIDFromLatLng(s2.LatLng{Lat: s1.Angle(p.X), Lng: s1.Angle(p.Y)}).Parent(level)
		return uint64(cellID)
	}
	return -1
}

func S2GetCenter(cellID uint64, level int, srid SRID) *Point {
	cellIDS2 := s2.CellID(cellID)
	if srid == SRID_WGS84_GPS {
		coord := cellIDS2.LatLng()
		return &Point{float64(coord.Lat), float64(coord.Lng)}
	}
	return nil
}
