package geo

import "github.com/hzw456/go-geo/geos"

// interface 对应geo中的geometry
func GeosToGeo(g *geos.CGeometry) Geometry {
	switch g.GetType() {
	case geos.POINT:
		return Point{g.GetCoord().X, g.GetCoord().Y}
	case geos.LINESTRING:
		return *NewLineString(exactPtsFromPoi(g)...)
	// case geos.LINEARRING:
	// 	var pts []Point
	// 	for _, v := range g.GetCoords() {
	// 		pts = append(pts, Point{v.X, v.Y})
	// 	}
	// 	return *NewLinearRing(pts...)
	case geos.MULTIPOINT:
		return *NewMultiPoint(exactPtsFromPoi(g)...)
	case geos.POLYGON:
		lines := exactLinesFromPoi(g)
		poly := Polygon{}
		for _, line := range lines {
			poly = append(poly, line.ToRing())
		}
		return poly
	case geos.MULTILINESTRING:
		lines := exactLinesFromPoi(g)
		return *NewMultiLineString(lines...)
	case geos.MULTIPOLYGON:
		polys := exactPolysFromPoi(g)
		return *NewMultiPolygon(polys...)
		// case geos.GEOMETRYCOLLECTION:
		// 	return "GeometryCollection", ELEM_COLLECTION
	}
	return nil
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

func exactPolysFromPoi(g *geos.CGeometry) (polys []Polygon) {
	for _, coord3D := range g.GetCoord3D() {
		poly := *NewPolygon()
		for i, coords := range coord3D {
			line := NewLineString()
			for _, coord := range coords {
				line.Append(Point{coord.X, coord.Y})
			}
			ring := line.ToRing()
			poly[i] = ring
		}
		polys = append(polys, poly)
	}
	return
}
