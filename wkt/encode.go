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

func GeometryToWktStringZ(x, y, z float64) string {
	return strconv.FormatFloat(x, 'f', -1, 64) + " " + strconv.FormatFloat(y, 'f', -1, 64) + " " + strconv.FormatFloat(z, 'f', -1, 64)
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

func PathZToWktString(pts []geo.PointZ) string {
	var res []string
	for _, pt := range pts {
		w := GeometryToWktStringZ(pt.X, pt.Y, pt.Z)
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

func PolygonZToWktString(lines []geo.LineStringZ) string {
	var res []string
	for _, line := range lines {
		w := PathZToWktString(line)
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

func PolygonZToWktString2(lines []geo.LinearRingZ) string {
	var res []string
	for _, line := range lines {
		w := PathZToWktString(line)
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

func MultiPolygonZToWktString(polys []geo.PolygonZ) string {
	var res []string
	for _, poly := range polys {
		w := PolygonZToWktString2(poly)
		res = append(res, string(w))
	}
	sj := strings.Join(res, ",")
	return "(" + sj + ")"
}

func Encode(geom geo.Geometry) string {
	switch geom := geom.(type) {
	case geo.Point:
		return fmt.Sprintf("POINT(%s)", GeometryToWktString(geom.X, geom.Y))
	case geo.PointZ:
		return fmt.Sprintf("POINT Z (%s)", GeometryToWktStringZ(geom.X, geom.Y, geom.Z))
	case geo.MultiPoint:
		return fmt.Sprintf("MULTIPOINT%s", PathToWktString(geom))
	case geo.MultiPointZ:
		return fmt.Sprintf("MULTIPOINT Z %s", PathZToWktString(geom))
	case geo.LineString:
		return fmt.Sprintf("LINESTRING%s", PathToWktString(geom))
	case geo.LineStringZ:
		return fmt.Sprintf("LINESTRING Z %s", PathZToWktString(geom))
	case geo.MultiLineString:
		return fmt.Sprintf("MULTILINESTRING%s", PolygonToWktString(geom))
	case geo.MultiLineStringZ:
		return fmt.Sprintf("MULTILINESTRING Z %s", PolygonZToWktString(geom))
	case geo.Polygon:
		return fmt.Sprintf("POLYGON%s", PolygonToWktString2(geom))
	case geo.PolygonZ:
		return fmt.Sprintf("POLYGON Z %s", PolygonZToWktString2(geom))
	case geo.MultiPolygon:
		return fmt.Sprintf("MULTIPOLYGON%s", MultiPolygonToWktString(geom))
	case geo.MultiPolygonZ:
		return fmt.Sprintf("MULTIPOLYGON Z %s", MultiPolygonZToWktString(geom))
	default:
		return ""
	}
}
