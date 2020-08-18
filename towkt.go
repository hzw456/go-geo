package geo

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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
