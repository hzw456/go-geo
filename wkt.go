package gogeo

import (
	"fmt"
)

func GeoToWkt(geo Geometry) string {
	switch geo := geo.(type) {
	case Polygon:
		return polygonToWkt(geo)
	case Point:
		return pointToWkt(geo)
	case MultiPoint:
		return pointToWkt(geo...)
	case LineString:
		return lineToWkt(geo)
	case MultiLineString:
		return lineToWkt(geo...)
	case MultiPolygon:
		return polygonToWkt(geo...)
	}
	return ""
}

func pointToWkt(points ...Point) (wkt string) {
	isMultipoint := false
	if len(points) != 1 {
		isMultipoint = true
	}
	if isMultipoint {
		wkt = "MULTIPOINT("
	} else {
		wkt = "POINT"
	}
	for k, v := range points {
		wkt = wkt + "("
		wkt = wkt + fmt.Sprint(v.X) + " " + fmt.Sprint(v.Y)
		wkt = wkt + ")"
		if isMultipoint && k != len(points)-1 {
			wkt = wkt + ","
		}
	}

	if isMultipoint {
		wkt = wkt + ")"
	}
	return
}

func lineToWkt(lines ...LineString) (wkt string) {
	isMultiline := false
	if len(lines) != 1 {
		isMultiline = true
	}
	if isMultiline {
		wkt = "MULTILINESTRING("
	} else {
		wkt = "LINESTRING"
	}
	for _, v := range lines {
		wkt = wkt + "("
		for kk, vv := range v {
			wkt = wkt + fmt.Sprint(vv.X) + " " + fmt.Sprint(vv.Y)
			if kk != len(v)-1 {
				wkt = wkt + ","
			}
		}
		wkt = wkt + ")"
	}
	if isMultiline {
		wkt = wkt + ")"
	}
	return
}

func polygonToWkt(polys ...Polygon) (wkt string) {
	isMultipoly := false
	if len(polys) != 1 {
		isMultipoly = true
	}
	if isMultipoly {
		wkt = "MULTIPOLYGON("
	} else {
		wkt = "POLYGON"
	}
	for k, poly := range polys {
		for _, v := range poly {
			wkt = wkt + "(("
			for kk, vv := range v {
				wkt = wkt + fmt.Sprint(vv.X) + " " + fmt.Sprint(vv.Y)
				if kk != len(v)-1 {
					wkt = wkt + ","
				}
			}
			wkt = wkt + "))"
			if isMultipoly && k != len(polys)-1 {
				wkt = wkt + ","
			}
		}
	}

	if isMultipoly {
		wkt = wkt + ")"
	}
	return
}
