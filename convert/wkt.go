package convert

import (
	"fmt"

	"git.xiaojukeji.com/haozhiwei/go-geo/element"
)

func PolygonToWkt(poly element.Polygon) (string, error) {
	var wkt string
	isMultipoly := false
	if len(poly) != 1 {
		isMultipoly = true
	}
	if isMultipoly {
		wkt = "MULTIPOLYGON("
	} else {
		wkt = "POLYGON"
	}
	for k, v := range poly {
		wkt = wkt + "(("
		for kk, vv := range v {
			wkt = wkt + fmt.Sprint(vv.X) + " " + fmt.Sprint(vv.Y)
			if kk != len(v)-1 {
				wkt = wkt + ","
			}
		}
		wkt = wkt + "))"
		if isMultipoly && k != len(poly)-1 {
			wkt = wkt + ","
		}
	}
	if isMultipoly {
		wkt = wkt + ")"
	}
	return wkt, nil
}
