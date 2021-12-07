package geo

import "math"

func Mbr(geom Geometry) Polygon {
	var coords []Point
	switch geom := geom.(type) {
	case Point:
		return Polygon{}
	case MultiPoint:
		coords = geom.GetPointSet()
	case LineString:
		coords = geom.GetPointSet()
	case MultiLineString:
		for _, l := range geom {
			coords = append(coords, l.GetPointSet()...)
		}
	case Polygon:
		coords = geom.GetExteriorPoints()
	case MultiPolygon:
		for _, p := range geom {
			coords = append(coords, p.GetExteriorPoints()...)
		}
	}
	convexHull := ConvexHull(coords...)
	if convexHull == nil {
		return nil
	}
	cpt := Centroid(convexHull)
	if math.IsNaN(cpt.X) || math.IsNaN(cpt.Y) {
		return nil
	}
	minArea := math.MaxFloat64
	minAngle := 0.0
	ci := coords[0]
	var ssr Geometry
	for i := 0; i < len(coords)-1; i++ {
		cii := coords[i+1]
		angle := math.Atan2(cii.Y-ci.Y, cii.X-ci.X)
		rect := BoundingBox(RotateCW(convexHull, cpt, angle))
		area := GetArea(BoxToGeo(*rect))
		if area < minArea {
			minArea = area
			ssr = BoxToGeo(*rect)
			minAngle = angle
		}
		ci = cii
	}
	polyGeo := RotateCCW(ssr, cpt, minAngle)
	if polyGeo == nil {
		return nil
	}
	return polyGeo.(Polygon)
}
