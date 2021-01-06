package geo

import (
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
)

func S2GetCellID(p Point) int64 {
	cellID := s2.CellIDFromLatLng(s2.LatLng{Lat: s1.Angle(p.X), Lng: s1.Angle(p.Y)})
	return int64(cellID)
}
