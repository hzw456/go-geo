package geo

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func PointToWkt(points ...Point) (wkt string) {
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

func LineToWkt(lines ...LineString) (wkt string) {
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

func PolygonToWkt(polys ...Polygon) (wkt string) {
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

func FromWkt(wkt string) (Geometry, error) {
	switch {
	case strings.HasPrefix(wkt, "POINT"):
		return WktToPoint(wkt)
	case strings.HasPrefix(wkt, "LINESTRING"):
		return WktToLineString(wkt)
	case strings.HasPrefix(wkt, "POLYGON"):
		return WktToPolygon(wkt)
		// case strings.HasPrefix(wkt, "MULTIPOINT"):
		// 	return
		// case strings.HasPrefix(wkt, "MULTILINESTRING"):
		// 	return
		// case strings.HasPrefix(wkt, "MULITIPOLYGON"):
		// 	return
		// case strings.HasPrefix(wkt, "GEOMETRYCOLLECTION"):
		// 	return
	}
	return nil, nil
}

func WktToPoint(str string) (Point, error) {
	pts, err := wktProcess(str)
	if err != nil || len(pts) != 1 {
		return Point{}, err
	}
	return pts[0], nil
}

func WktToLineString(str string) (LineString, error) {
	pts, err := wktProcess(str)
	if err != nil || len(pts) < 2 {
		return LineString{}, err
	}
	return *NewLineString(pts...), nil
}

func WktToPolygon(str string) (Polygon, error) {
	pts, err := wktProcess(str)
	if err != nil || len(pts) != 1 {
		return Polygon{}, err
	}
	return *NewPolygonFromPois(pts...), nil
}

func wktProcess(str string) ([]Point, error) {
	re, _ := regexp.Compile(`-?(?:\.\d+|\d+(?:\.\d*)?) -?(?:\.\d+|\d+(?:\.\d*)?)`)
	all := re.FindAllString(str, -1)
	var pts []Point
	for _, item := range all {
		fmt.Println(string(item))
		strs := strings.Split(item, " ")
		if len(strs) != 2 {
			return nil, errors.New("invalid wkt string")
		}
		x, err := strconv.ParseFloat(strs[1], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(strs[0], 64)
		if err != nil {
			return nil, err
		}
		pt := NewPoint(x, y)
		pts = append(pts, *pt)
	}
	return pts, nil
}
