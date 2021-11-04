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

func S2GetBoundary(cellID uint64, srid SRID) []Point {
	cellIDS2 := s2.CellID(cellID)
	var res []Point
	if srid == SRID_WGS84_GPS {
		loop := s2.LoopFromCell(s2.CellFromCellID(cellIDS2))
		for _, v := range loop.Vertices() {
			res = append(res, Point{s2.LatLngFromPoint(v).Lat.Degrees(), s2.LatLngFromPoint(v).Lng.Degrees()})
		}
		return res
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

func RegionCoverer(poly Polygon, level int, srid SRID) []uint64 {
	if !poly.IsCCW() {
		pois := poly.GetExteriorRing().GetPointSet()
		for i, j := 0, len(pois)-1; i < j; i, j = i+1, j-1 {
			pois[i], pois[j] = pois[j], pois[i]
		}
		poly = *NewPolygonFromPois(pois...)
	}
	loops := []*s2.Loop{}
	loops = append(loops, toLoop(poly.GetExteriorPoints()))
	var cellIDsInt []uint64
	polygon := s2.PolygonFromLoops(loops)
	latlng := s2.LatLngFromDegrees(poly.GetExteriorPoints()[0].X, poly.GetExteriorPoints()[0].Y)
	cellIDs := s2.SimpleRegionCovering(s2.Region(polygon), s2.PointFromLatLng(latlng), level)
	for _, cellID := range cellIDs {
		cellIDsInt = append(cellIDsInt, uint64(cellID))
	}

	return cellIDsInt
}

func toLoop(points []Point) *s2.Loop {
	var pts []s2.Point
	for _, pt := range points {
		pts = append(pts, s2.PointFromLatLng(s2.LatLngFromDegrees(pt.X, pt.Y)))
	}
	for i, j := 0, len(pts)-1; i < j; i, j = i+1, j-1 {
		pts[i], pts[j] = pts[j], pts[i]
	}
	return s2.LoopFromPoints(pts)
}
