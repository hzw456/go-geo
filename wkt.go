package geo

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func (p Point) ToWkt() ([]byte, error) {
	return []byte(fmt.Sprintf("POINT(%s)", GeometryToWktString(p.X, p.Y))), nil
}

func (mp MultiPoint) ToWkt() ([]byte, error) {
	return []byte(fmt.Sprintf("MULTIPOINT%s", PathToWktString(mp))), nil
}

func (l LineString) ToWkt() ([]byte, error) {
	return []byte(fmt.Sprintf("LINESTRING%s", PathToWktString(l))), nil
}

func (ml MultiLineString) ToWkt() ([]byte, error) {
	return []byte(fmt.Sprintf("MULTILINESTRING%s", PolygonToWktString(ml))), nil
}

func (p Polygon) ToWkt() ([]byte, error) {
	return []byte(fmt.Sprintf("POLYGON%s", PolygonToWktString2(p))), nil
}

func (mp MultiPolygon) ToWkt() ([]byte, error) {
	return []byte(fmt.Sprintf("MULTIPOLYGON%s", MultiPolygonToWktString(mp))), nil
}

func GeometryToWktString(x, y float64) string {
	return strconv.FormatFloat(x, 'f', -1, 64) + " " + strconv.FormatFloat(y, 'f', -1, 64)
}

func PathToWktString(pts []Point) string {
	var res []string
	for _, pt := range pts {
		w := GeometryToWktString(pt.X, pt.Y)
		res = append(res, string(w))
	}
	sj := strings.Join(res, ",")
	return "(" + sj + ")"
}

func PolygonToWktString(lines []LineString) string {
	var res []string
	for _, line := range lines {
		w := PathToWktString(line)
		res = append(res, string(w))
	}
	sj := strings.Join(res, ",")
	return "(" + sj + ")"
}

func PolygonToWktString2(lines []LinearRing) string {
	var res []string
	for _, line := range lines {
		w := PathToWktString(line)
		res = append(res, string(w))
	}
	sj := strings.Join(res, ",")
	return "(" + sj + ")"
}

func MultiPolygonToWktString(polys []Polygon) string {
	var res []string
	for _, poly := range polys {
		w := PolygonToWktString2(poly)
		res = append(res, string(w))
	}
	sj := strings.Join(res, ",")
	return "(" + sj + ")"
}

func FromWkt(wkt string) (Geometry, error) {
	switch {
	case strings.HasPrefix(wkt, "POINT"):
		return PointFromWKT(wkt)
	case strings.HasPrefix(wkt, "LINESTRING"):
		return LineStringFromWKT(wkt)
	case strings.HasPrefix(wkt, "POLYGON"):
		return PolygonFromWKT(wkt)
	case strings.HasPrefix(wkt, "MULTIPOINT"):
		return MultiPointFromWKT(wkt)
	case strings.HasPrefix(wkt, "MULTILINESTRING"):
		return MultiLineStringFromWKT(wkt)
	case strings.HasPrefix(wkt, "MULITIPOLYGON"):
		return MultiPointFromWKT(wkt)
		// case strings.HasPrefix(wkt, "GEOMETRYCOLLECTION"):
		// 	return
	}
	return nil, nil
}

func PointFromWKT(wkt string) (Point, error) {
	strs := strings.Split(strings.Trim(wkt, " "), " ")
	x, err := strconv.ParseFloat(strs[1], 64)
	if err != nil {
		return Point{}, err
	}
	y, err := strconv.ParseFloat(strs[0], 64)
	if err != nil {
		return Point{}, err
	}
	return Point{x, y}, nil
}

func MultiPointFromWKT(wkt string) (MultiPoint, error) {
	wkt = strings.TrimLeft(wkt, "(")
	wkt = strings.TrimRight(wkt, ")")
	wkt = strings.Trim(wkt, " ")
	terms := strings.Split(wkt, ",")
	if len(terms) != 2 {
		return nil, errors.New("invalid wkt string")
	}
	var mp MultiPoint
	for _, term := range terms {
		strs := strings.Split(strings.TrimRight(strings.TrimLeft(term, "("), ")"), " ")
		x, err := strconv.ParseFloat(strs[1], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(strs[0], 64)
		if err != nil {
			return nil, err
		}
		mp.Append(Point{x, y})
	}
	return mp, nil
}

func LineStringFromWKT(wkt string) (LineString, error) {
	wkt = strings.TrimLeft(wkt, "(")
	wkt = strings.TrimRight(wkt, ")")
	terms := strings.Split(wkt, ",")
	var linestring LineString
	for _, term := range terms {
		strs := strings.Split(strings.Trim(term, " "), " ")
		// if prevgeopos != null {
		// 	continue
		// }
		x, err := strconv.ParseFloat(strs[1], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(strs[0], 64)
		if err != nil {
			return nil, err
		}
		linestring.Append(Point{x, y})
	}
	return linestring, nil
}

func MultiLineStringFromWKT(wkt string) (MultiLineString, error) {
	re := regexp.MustCompile(`(\ *[(]\ *(?:\ *(?:[0-9-.Ee]+[ ]+[0-9-.Ee]+)[, ]*\ *)*\ *[)])[, ]*`)
	matches := re.FindStringSubmatch(wkt)
	var ml MultiLineString
	for _, v := range matches {
		linestring, err := LineStringFromWKT(v)
		if err != nil {
			return nil, err
		}
		ml = append(ml, linestring)
	}

	return ml, nil
}

func PolygonFromWKT(wkt string) (Polygon, error) {
	re := regexp.MustCompile(`(\ *[(]\ *(?:\ *(?:[0-9-.Ee]+[ ]+[0-9-.Ee]+)[, ]*\ *)*\ *[)])[, ]*`)
	matches := re.FindStringSubmatch(wkt)
	var poly Polygon
	for _, v := range matches {
		linestring, err := LineStringFromWKT(v)
		if err != nil {
			return nil, err
		}
		poly = append(poly, linestring.ToRing())
	}
	return poly, nil
}

func MultiPolygonFromWKT(wkt string) (MultiPolygon, error) {
	re := regexp.MustCompile(`([(](?:\ *[(]\ *(?:\ *(?:[0-9-.]+[ ]+[0-9-.]+)[, ]*\ *)*\ *[)][, ]*)*[)])`)
	matches := re.FindStringSubmatch(wkt)
	var multiPoly MultiPolygon
	for _, v := range matches {
		polygon, err := PolygonFromWKT(v)
		if err != nil {
			return nil, err
		}
		multiPoly = append(multiPoly, polygon)
	}
	return multiPoly, nil
}

func WktToGeometry(wkt string) (Geometry, error) {
	wkt = strings.Trim(wkt, " ")
	wkt = strings.Replace(wkt, ", ", ",", -1)
	re := regexp.MustCompile(`([A-Z]+)\s*[(]\s*(\(*.+\)*)\s*[)]`)
	match := re.FindStringSubmatch(wkt)

	switch match[1] {
	case "MULTIPOLYGON":
		return MultiPolygonFromWKT(match[2])
	case "POLYGON":
		return PolygonFromWKT(match[2])
	case "MULTILINESTRING":
		return MultiLineStringFromWKT(match[2])
	case "LINESTRING":
		return LineStringFromWKT(match[2])
	case "MULTIPOINT":
		return MultiPointFromWKT(match[2])
	case "POINT":
		return PointFromWKT(match[2])
	}
	return nil, errors.New("wkt format is error")
}
