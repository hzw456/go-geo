package wkt

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hzw456/go-geo"
)

func GeometryToWktString(x, y float64) string {
	return strconv.FormatFloat(x, 'f', -1, 64) + " " + strconv.FormatFloat(y, 'f', -1, 64)
}

func PathToWktString(pts []geo.Point) string {
	var res []string
	for _, pt := range pts {
		w := GeometryToWktString(pt.X, pt.Y)
		res = append(res, string(w))
	}
	sj := strings.Join(res, ",")
	return "(" + sj + ")"
}

func PolygonToWktString(lines []geo.LineString) string {
	var res []string
	for _, line := range lines {
		w := PathToWktString(line)
		res = append(res, string(w))
	}
	sj := strings.Join(res, ",")
	return "(" + sj + ")"
}

func PolygonToWktString2(lines []geo.LinearRing) string {
	var res []string
	for _, line := range lines {
		w := PathToWktString(line)
		res = append(res, string(w))
	}
	sj := strings.Join(res, ",")
	return "(" + sj + ")"
}

func MultiPolygonToWktString(polys []geo.Polygon) string {
	var res []string
	for _, poly := range polys {
		w := PolygonToWktString2(poly)
		res = append(res, string(w))
	}
	sj := strings.Join(res, ",")
	return "(" + sj + ")"
}

func Encode(geom geo.Geometry) string {
	switch geom := geom.(type) {
	case geo.Point:
		return fmt.Sprintf("POINT(%s)", GeometryToWktString(geom.X, geom.Y))
	case geo.MultiPoint:
		return fmt.Sprintf("MULTIPOINT%s", PathToWktString(geom))
	case geo.LineString:
		return fmt.Sprintf("LINESTRING%s", PathToWktString(geom))
	case geo.MultiLineString:
		return fmt.Sprintf("MULTILINESTRING%s", PolygonToWktString(geom))
	case geo.Polygon:
		return fmt.Sprintf("POLYGON%s", PolygonToWktString2(geom))
	case geo.MultiPolygon:
		return fmt.Sprintf("MULTIPOLYGON%s", MultiPolygonToWktString(geom))
	default:
		return ""
	}
}
