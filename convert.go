package geo

import "github.com/sadnessly/go-geo/geos"

// interface 对应geo中的geometry
func GeosToGeo(g *geos.CGeometry) (Geometry, GeometryType) {
	switch g.GetType() {
	case geos.POINT:
		return Point{g.GetCoord().X, g.GetCoord().Y}, ELEM_POINT
	case geos.LINESTRING:
		return *NewLineString(exactPtsFromPoi(g)...), ELEM_LINESTRING
	// case geos.LINEARRING:
	// 	var pts []Point
	// 	for _, v := range g.GetCoords() {
	// 		pts = append(pts, Point{v.X, v.Y})
	// 	}
	// 	return *NewLinearRing(pts...)
	case geos.MULTIPOINT:
		return *NewMultiPoint(exactPtsFromPoi(g)...), ELEM_MULTIPOINT
	case geos.POLYGON:
		lines := exactLinesFromPoi(g)
		poly := Polygon{}
		for _, line := range lines {
			poly = append(poly, line.ToRing())
		}
		return poly, ELEM_POLYGON
	case geos.MULTILINESTRING:
		lines := exactLinesFromPoi(g)
		return *NewMultiLineString(lines...), ELEM_MULTILINESTRING
		// case geos.MULTIPOLYGON:
		// 	return "MultiPolygon", ELEM_MULTIPOLYGON
		// case geos.GEOMETRYCOLLECTION:
		// 	return "GeometryCollection", ELEM_COLLECTION
	}
	return nil, ELEM_UNKNOWN
}

func exactPtsFromPoi(g *geos.CGeometry) (pts []Point) {
	for _, v := range g.GetCoords() {
		pts = append(pts, Point{v.X, v.Y})
	}
	return
}

func exactLinesFromPoi(g *geos.CGeometry) (lines []LineString) {
	for _, coords := range g.GetCoordsSlice() {
		ls := NewLineString()
		for _, coord := range coords {
			ls.Append(Point{coord.X, coord.Y})
		}
		lines = append(lines, *ls)
	}
	return
}
