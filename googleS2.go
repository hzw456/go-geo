package geo

import (
	"github.com/golang/geo/s2"
)

func S2GetCellID(p Point, level int, srid SRID) uint64 {
	if srid == SRID_WGS84_GPS {
		cellID := s2.CellIDFromLatLng(s2.LatLngFromDegrees(p.X, p.Y)).Parent(level)
		return uint64(cellID)
	}
	return 0
}

func S2GetCenter(cellID uint64, srid SRID) *Point {
	cellIDS2 := s2.CellID(cellID)
	if srid == SRID_WGS84_GPS {
		coord := cellIDS2.LatLng()
		return &Point{coord.Lat.Degrees(), coord.Lng.Degrees()}
	}
	return nil
}

func S2GetAllNeighbors(cellID uint64, srid SRID) []uint64 {
	cellIDS2 := s2.CellID(cellID)
	var cellIDsInt []uint64
	if srid == SRID_WGS84_GPS {
		cellIDs := cellIDS2.AllNeighbors(cellIDS2.Level())
		for _, cellID := range cellIDs {
			cellIDsInt = append(cellIDsInt, uint64(cellID))
		}

		return cellIDsInt
	}
	return []uint64{}
}
