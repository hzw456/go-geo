package geojson

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hzw456/go-geo"
)

type GeoJson struct {
	Type        string                 `json:"type"`
	Coordinates interface{}            `json:"coordinates,omitempty"`
	Geometries  interface{}            `json:"geometries,omitempty"`
	CRS         map[string]interface{} `json:"crs,omitempty"`
}

func MarshalGeo(geom geo.Geometry) ([]byte, error) {
	var geoj GeoJson
	var err error
	switch geom := geom.(type) {
	case geo.Point:
		geoj.Type = "Point"
		geoj.Coordinates, err = ConvertPoi(geom)
		if err != nil {
			return nil, err
		}
	case geo.PointZ:
		geoj.Type = "Point"
		geoj.Coordinates, err = ConvertPoiZ(geom)
		if err != nil {
			return nil, err
		}
	case geo.MultiPoint:
		geoj.Type = "MultiPoint"
		geoj.Coordinates, err = ConvertPoiSet(geom)
		if err != nil {
			return nil, err
		}
	case geo.MultiPointZ:
		geoj.Type = "MultiPoint"
		geoj.Coordinates, err = ConvertPoiZSet(geom)
		if err != nil {
			return nil, err
		}
	case geo.LineString:
		geoj.Type = "LineString"
		geoj.Coordinates, err = ConvertPoiSet(geom)
		if err != nil {
			return nil, err
		}
	case geo.LineStringZ:
		geoj.Type = "LineString"
		geoj.Coordinates, err = ConvertPoiZSet(geom)
		if err != nil {
			return nil, err
		}
	case geo.MultiLineString:
		geoj.Type = "MultiLineString"
		geoj.Coordinates, err = ConvertPathSet(geom)
		if err != nil {
			return nil, err
		}
	case geo.MultiLineStringZ:
		geoj.Type = "MultiLineString"
		geoj.Coordinates, err = ConvertPathZSet(geom)
		if err != nil {
			return nil, err
		}
	case geo.Polygon:
		geoj.Type = "Polygon"
		geoj.Coordinates, err = ConvertPathSet(geom)
		if err != nil {
			return nil, err
		}
	case geo.PolygonZ:
		geoj.Type = "Polygon"
		geoj.Coordinates, err = ConvertPathSet(geom)
		if err != nil {
			return nil, err
		}
	case geo.MultiPolygon:
		geoj.Type = "MultiPolygon"
		geoj.Coordinates, err = ConvertPolygonSet(geom)
		if err != nil {
			return nil, err
		}
	case geo.MultiPolygonZ:
		geoj.Type = "MultiPolygon"
		geoj.Coordinates, err = ConvertPolygonZSet(geom)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("no such type")
	}
	return json.Marshal(geoj)
}

func UnmarshalGeo(gjson *GeoJson) (geo.Geometry, error) {
	switch gjson.Type {
	case "Point":
		isZ := isInterPointZ(gjson.Coordinates)
		if isZ {
			return inter2PointZ(gjson.Coordinates), nil
		}
		return inter2Point(gjson.Coordinates), nil
	case "MultiPoint":
		return nil, errors.New("Not Support MultiPoint (TODO)")
	case "LineString":
		isZ := isInterLineZ(gjson.Coordinates)
		if isZ {
			return inter2LineStringZ(gjson.Coordinates), nil
		}
		return inter2LineString(gjson.Coordinates), nil
	case "MultiLineString":
		return nil, errors.New("Not Support MultiLineString (TODO)")
	case "Polygon":
		isZ := isInterPolygonZ(gjson.Coordinates)
		if isZ {
			return inter2PolygonZ(gjson.Coordinates), nil
		}
		return inter2Polygon(gjson.Coordinates), nil
	case "MultiPolygon":
		return nil, errors.New("Not Support MultiPolygon (TODO)")
	}
	return nil, nil
}

// func (c Collection) ToGeojson() ([]byte, error) {
// 	return "GeometryCollection"
// }

func ConvertPoi(p geo.Point) ([]float64, error) {
	return []float64{p.X, p.Y}, nil
}

func ConvertPoiZ(p geo.PointZ) ([]float64, error) {
	return []float64{p.X, p.Y, p.Z}, nil
}

// only linestring and multipoint
func ConvertPoiSet(g geo.Geometry) ([][]float64, error) {
	var pts []geo.Point
	switch g := g.(type) {
	case geo.LineString:
		pts = g.GetPointSet()
	case geo.MultiPoint:
		pts = g.GetPointSet()
	default:
		return nil, errors.New("could not parsing geometry besides linesting and multipoint")
	}
	var interface2D [][]float64
	for _, pt := range pts {
		res, err := ConvertPoi(pt)
		if err != nil {
			return nil, err
		}
		interface2D = append(interface2D, res)
	}
	return interface2D, nil
}

// only linestring and multipoint
func ConvertPoiZSet(g geo.Geometry) ([][]float64, error) {
	var pts []geo.PointZ
	switch g := g.(type) {
	case geo.LineStringZ:
		pts = g.GetPointSet()
	case geo.MultiPointZ:
		pts = g.GetPointSet()
	default:
		return nil, errors.New("could not parsing geometry besides linesting and multipoint")
	}
	var interface2D [][]float64
	for _, pt := range pts {
		res, err := ConvertPoiZ(pt)
		if err != nil {
			return nil, err
		}
		interface2D = append(interface2D, res)
	}
	return interface2D, nil
}

// only multilinestring and polygon
func ConvertPathSet(g geo.Geometry) ([][][]float64, error) {
	var lines []geo.LineString
	switch g := g.(type) {
	case geo.MultiLineString:
		for _, line := range g {
			lines = append(lines, line)
		}
	case geo.Polygon:
		for _, ring := range g {
			lines = append(lines, ring.ToLineString())
		}
	default:
		return nil, errors.New("could not parsing geometry besides multilinestring and polygon")
	}
	var interface3D [][][]float64
	for _, line := range lines {
		res, err := ConvertPoiSet(line)
		if err != nil {
			return nil, err
		}
		interface3D = append(interface3D, res)
	}
	return interface3D, nil
}

// only multilinestring and polygon
func ConvertPathZSet(g geo.Geometry) ([][][]float64, error) {
	var lines []geo.LineStringZ
	switch g := g.(type) {
	case geo.MultiLineStringZ:
		for _, line := range g {
			lines = append(lines, line)
		}
	case geo.PolygonZ:
		for _, ring := range g {
			lines = append(lines, ring.ToLineString())
		}
	default:
		return nil, errors.New("could not parsing geometry besides multilinestring and polygon")
	}
	var interface3D [][][]float64
	for _, line := range lines {
		res, err := ConvertPoiZSet(line)
		if err != nil {
			return nil, err
		}
		interface3D = append(interface3D, res)
	}
	return interface3D, nil
}

//only multipolygon
func ConvertPolygonSet(g geo.Geometry) ([][][][]float64, error) {
	var polys []geo.Polygon
	switch g := g.(type) {
	case geo.MultiPolygon:
		for _, poly := range g {
			polys = append(polys, poly)
		}
	default:
		return nil, errors.New("could not parsing geometry besides multipolygon")
	}
	var interface4D [][][][]float64
	for _, poly := range polys {
		res, err := ConvertPathSet(poly)
		if err != nil {
			return nil, err
		}
		interface4D = append(interface4D, res)
	}
	return interface4D, nil
}

//only multipolygon
func ConvertPolygonZSet(g geo.Geometry) ([][][][]float64, error) {
	var polys []geo.PolygonZ
	switch g := g.(type) {
	case geo.MultiPolygonZ:
		for _, poly := range g {
			polys = append(polys, poly)
		}
	default:
		return nil, errors.New("could not parsing geometry besides multipolygon")
	}
	var interface4D [][][][]float64
	for _, poly := range polys {
		res, err := ConvertPathZSet(poly)
		if err != nil {
			return nil, err
		}
		interface4D = append(interface4D, res)
	}
	return interface4D, nil
}

func decodePosition(data interface{}) ([]float64, error) {
	coords, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("not a valid position, got %v", data)
	}

	result := make([]float64, 0, len(coords))
	for _, coord := range coords {
		if f, ok := coord.(float64); ok {
			result = append(result, f)
		} else {
			return nil, fmt.Errorf("not a valid coordinate, got %v", coord)
		}
	}

	return result, nil
}

func decodePositionSet(data interface{}) ([][]float64, error) {
	points, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("not a valid set of positions, got %v", data)
	}

	result := make([][]float64, 0, len(points))
	for _, point := range points {
		if p, err := decodePosition(point); err == nil {
			result = append(result, p)
		} else {
			return nil, err
		}
	}

	return result, nil
}

func decodePathSet(data interface{}) ([][][]float64, error) {
	sets, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("not a valid path, got %v", data)
	}

	result := make([][][]float64, 0, len(sets))

	for _, set := range sets {
		if s, err := decodePositionSet(set); err == nil {
			result = append(result, s)
		} else {
			return nil, err
		}
	}

	return result, nil
}

func decodePolygonSet(data interface{}) ([][][][]float64, error) {
	polygons, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("not a valid polygon, got %v", data)
	}

	result := make([][][][]float64, 0, len(polygons))
	for _, polygon := range polygons {
		if p, err := decodePathSet(polygon); err == nil {
			result = append(result, p)
		} else {
			return nil, err
		}
	}

	return result, nil
}

func interface2D(data interface{}) []interface{} {
	return data.([]interface{})
}

func inter2Point(data interface{}) geo.Point {
	var point geo.Point
	vI := interface2D(data)
	point.X = vI[0].(float64)
	point.Y = vI[1].(float64)
	return point
}

func inter2PointZ(data interface{}) geo.PointZ {
	var pointZ geo.PointZ
	vI := interface2D(data)
	pointZ.X = vI[0].(float64)
	pointZ.Y = vI[1].(float64)
	pointZ.Z = vI[2].(float64)
	return pointZ
}

func inter2LineString(data interface{}) geo.LineString {
	var lineString geo.LineString
	for _, v := range interface2D(data) {
		point := inter2Point(v)
		lineString = append(lineString, point)
	}
	return lineString
}

func inter2LineStringZ(data interface{}) geo.LineStringZ {
	var lineStringZ geo.LineStringZ
	for _, v := range interface2D(data) {
		pointZ := inter2PointZ(v)
		lineStringZ = append(lineStringZ, pointZ)
	}
	return lineStringZ
}

func inter2LinearRing(data interface{}) geo.LinearRing {
	var linearRing geo.LinearRing
	for _, v := range interface2D(data) {
		point := inter2Point(v)
		linearRing = append(linearRing, point)
	}
	return linearRing
}

func inter2LinearRingZ(data interface{}) geo.LinearRingZ {
	var linearRingZ geo.LinearRingZ
	for _, v := range interface2D(data) {
		pointZ := inter2PointZ(v)
		linearRingZ = append(linearRingZ, pointZ)
	}
	return linearRingZ
}

func inter2Polygon(data interface{}) geo.Polygon {
	var polygon geo.Polygon
	for _, line := range interface2D(data) {
		linearRing := inter2LinearRing(line)
		polygon = append(polygon, linearRing)
	}
	return polygon
}

func inter2PolygonZ(data interface{}) geo.PolygonZ {
	var polygonZ geo.PolygonZ
	for _, line := range interface2D(data) {
		linearRingZ := inter2LinearRingZ(line)
		polygonZ = append(polygonZ, linearRingZ)
	}
	return polygonZ
}

// 判断坐标是否有Z值
func isInterPointZ(data interface{}) bool {
	vI := interface2D(data)
	if len(vI) == 3 {
		return true
	}
	return false
}

// 判断坐标是否有Z值
func isInterLineZ(data interface{}) bool {
	for _, points := range interface2D(data) {
		vI := interface2D(points)
		if len(vI) == 3 {
			return true
		}
	}
	return false
}

// 判断坐标是否有Z值
func isInterPolygonZ(data interface{}) bool {
	for _, line := range interface2D(data) {
		for _, v := range interface2D(line) {
			vI := interface2D(v)
			if len(vI) == 3 {
				return true
			}
		}
	}
	return false
}
