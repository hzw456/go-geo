package wkt

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/hzw456/go-geo"
)

func PointFromWKT(wkt string) (geo.Geometry, error) {
	strs := strings.Split(strings.Trim(wkt, " "), " ")
	x, err := strconv.ParseFloat(strs[0], 64)
	if err != nil {
		return geo.Point{}, err
	}
	y, err := strconv.ParseFloat(strs[1], 64)
	if err != nil {
		return geo.Point{}, err
	}
	if len(strs) == 3 {
		z, err := strconv.ParseFloat(strs[2], 64)
		if err != nil {
			return geo.Point{}, err
		}
		return geo.PointZ{x, y, z}, nil
	}
	return geo.Point{x, y}, nil
}

func MultiPointFromWKT(wkt string) (geo.Geometry, error) {
	wkt = strings.TrimLeft(wkt, "(")
	wkt = strings.TrimRight(wkt, ")")
	wkt = strings.Trim(wkt, " ")
	terms := strings.Split(wkt, ",")
	var mp geo.MultiPoint
	var mpz geo.MultiPointZ
	for _, term := range terms {
		strs := strings.Split(strings.TrimRight(strings.TrimLeft(term, "("), ")"), " ")
		x, err := strconv.ParseFloat(strs[0], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(strs[1], 64)
		if err != nil {
			return nil, err
		}
		if len(strs) == 3 {
			z, err := strconv.ParseFloat(strs[2], 64)
			if err != nil {
				return geo.Point{}, err
			}
			mpz.Append(geo.PointZ{x, y, z})
		} else {
			mp.Append(geo.Point{x, y})
		}
	}
	if len(mp) != 0 {
		return mp, nil
	} else if len(mpz) != 0 {
		return mpz, nil
	} else {
		return mp, errors.New("multiPolygon is empty")
	}
}

func LineStringFromWKT(wkt string) (geo.Geometry, error) {
	wkt = strings.TrimLeft(wkt, "(")
	wkt = strings.TrimRight(wkt, ")")
	terms := strings.Split(wkt, ",")
	var linestring geo.LineString
	var linestringZ geo.LineStringZ
	for _, term := range terms {
		strs := strings.Split(strings.Trim(term, " "), " ")
		x, err := strconv.ParseFloat(strs[0], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(strs[1], 64)
		if err != nil {
			return nil, err
		}
		if len(strs) == 3 {
			z, err := strconv.ParseFloat(strs[2], 64)
			if err != nil {
				return geo.LineStringZ{}, err
			}
			linestringZ.Append(geo.PointZ{x, y, z})
		} else if len(strs) == 2 {
			linestring.Append(geo.Point{x, y})
		}
	}
	if len(linestring) > 0 {
		return linestring, nil
	} else if len(linestringZ) > 0 {
		return linestringZ, nil
	} else {
		return linestring, errors.New("linestring is empty")
	}

}

func MultiLineStringFromWKT(wkt string) (geo.Geometry, error) {
	re := regexp.MustCompile(`(\ *[(]\ *(?:\ *(?:[0-9-.Ee]+[ ]+[0-9-.Ee]+)[, ]*\ *)*\ *[)])[, ]*`)
	matches := re.FindStringSubmatch(wkt)
	var ml geo.MultiLineString
	var mlz geo.MultiLineStringZ
	for _, v := range matches {
		linestring, err := LineStringFromWKT(v)
		if err != nil {
			return nil, err
		}
		if linestring.Type() == geo.GEOMETRY_LINESTRING {
			line := linestring.(geo.LineString)
			ml = append(ml, line)
		} else if linestring.Type() == geo.GEOMETRY_LINESTRINGZ {
			lineZ := linestring.(geo.LineStringZ)
			mlz = append(mlz, lineZ)
		}
	}
	if len(ml) > 0 {
		return ml, nil
	} else if len(mlz) > 0 {
		return mlz, nil
	} else {
		return ml, errors.New("multilineString is empty")
	}
}

func PolygonFromWKT(wkt string) (geo.Geometry, error) {
	re := regexp.MustCompile(`^\(((\s*(\-|\+)?\d+(\.\d+)?\s+(\-|\+)?\d+(\.\d+)?\s+(\-|\+)?\d+(\.\d+)?),*)*\),*$`)
	matches := re.FindStringSubmatch(wkt)
	var poly geo.Polygon
	var polyz geo.PolygonZ
	for _, v := range matches {
		linestring, err := LineStringFromWKT(v)
		if err != nil {
			return nil, err
		}
		if linestring.Type() == geo.GEOMETRY_LINESTRING {
			line := linestring.(geo.LineString)
			poly = append(poly, line.ToRing())
		} else if linestring.Type() == geo.GEOMETRY_LINESTRINGZ {
			line := linestring.(geo.LineStringZ)
			polyz = append(polyz, line.ToRing())
		}
	}
	if len(poly) > 0 {
		return poly, nil
	} else if len(polyz) > 0 {
		return polyz, nil
	} else {
		return poly, errors.New("polygon is empty")
	}
}

func PolygonZFromWKT(wkt string) (geo.Geometry, error) {
	// re := regexp.MustCompile(`(\ *[(]\ *(?:\ *(?:[0-9-.Ee]+[ ]+[0-9-.Ee]+[ ]+[0-9-.Ee]+)[, ]*\ *)*\ *[)])[, ]*`)
	re := regexp.MustCompile(`^\(((\s*(\-|\+)?\d+(\.\d+)?\s+(\-|\+)?\d+(\.\d+)?\s+(\-|\+)?\d+(\.\d+)?),*)*\),*$`)
	// matches := re.FindStringSubmatch(wkt)
	matches := re.FindAllString(wkt, -1)
	var Poly geo.Polygon
	var PolyZ geo.PolygonZ
	for _, v := range matches {
		polygon, err := LineStringFromWKT(v)
		if err != nil {
			return nil, err
		}
		if polygon.Type() == geo.GEOMETRY_LINESTRING {
			line := polygon.(geo.LineString)
			Poly = append(Poly, line.ToRing())
		} else if polygon.Type() == geo.GEOMETRY_LINESTRINGZ {
			lineZ := polygon.(geo.LineStringZ)
			PolyZ = append(PolyZ, lineZ.ToRing())
		}

	}
	if len(Poly) > 0 {
		return Poly, nil
	} else if len(PolyZ) > 0 {
		return PolyZ, nil
	} else {
		return PolyZ, errors.New("polygon is empty")
	}

}

func MultiPolygonFromWKT(wkt string) (geo.Geometry, error) {
	re := regexp.MustCompile(`([(](?:\ *[(]\ *(?:\ *(?:[0-9-.]+[ ]+[0-9-.]+)[, ]*\ *)*\ *[)][, ]*)*[)])`)
	matches := re.FindStringSubmatch(wkt)
	var multiPoly geo.MultiPolygon
	var multiPolyZ geo.MultiPolygonZ
	for _, v := range matches {
		polygon, err := PolygonFromWKT(v)
		if err != nil {
			return nil, err
		}
		if polygon.Type() == geo.GEOMETRY_POLYGON {
			poly := polygon.(geo.Polygon)
			multiPoly = append(multiPoly, poly)
		} else if polygon.Type() == geo.GEOMETRY_POLYGONZ {
			poly := polygon.(geo.PolygonZ)
			multiPolyZ = append(multiPolyZ, poly)
		}

	}
	if len(multiPoly) > 0 {
		return multiPoly, nil
	} else if len(multiPolyZ) > 0 {
		return multiPolyZ, nil
	} else {
		return multiPoly, errors.New("polygon is empty")
	}
}

func Decode(wkt string) (geo.Geometry, error) {
	wkt = strings.Trim(wkt, " ")
	wkt = strings.Replace(wkt, " Z", "Z", 1)
	wkt = strings.Replace(wkt, ", ", ",", -1)
	re := regexp.MustCompile(`([A-Z]+)\s*[(]\s*(\(*.+\)*)\s*[)]`)

	match := re.FindStringSubmatch(wkt)
	if len(match) != 3 {
		return nil, errors.New("wkt format is error")
	}
	switch match[1] {
	case "MULTIPOLYGON", "MULTIPOLYGONZ":
		return MultiPolygonFromWKT(match[2])
	case "POLYGON":
		return PolygonFromWKT(match[2])
	case "MULTILINESTRING", "MULTILINESTRINGZ":
		return MultiLineStringFromWKT(match[2])
	case "LINESTRING", "LINESTRINGZ":
		return LineStringFromWKT(match[2])
	case "MULTIPOINT", "MULTIPOINTZ":
		return MultiPointFromWKT(match[2])
	case "POINT", "POINTZ":
		return PointFromWKT(match[2])
	case "POLYGONZ":
		return PolygonZFromWKT(match[2])
	}
	return nil, errors.New("wkt format is error")
}
