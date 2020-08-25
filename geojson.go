package geo

import (
	"encoding/json"
	"errors"
	"fmt"
)

type GeoJson struct {
	Type        string                 `json:"type"`
	Coordinates interface{}            `json:"coordinates,omitempty"`
	Geometries  interface{}            `json:"geometries,omitempty"`
	CRS         map[string]interface{} `json:"crs,omitempty"`
}

func (p Point) ToGeojson() ([]byte, error) {
	var geoj GeoJson
	geoj.Type = p.Type()
	geom, err := ConvertPoi(p)
	if err != nil {
		return nil, err
	}
	geoj.Coordinates = geom
	return json.Marshal(geoj)
}

func (mp MultiPoint) ToGeojson() ([]byte, error) {
	var geoj GeoJson
	geoj.Type = mp.Type()
	geom, err := ConvertPoiSet(mp)
	if err != nil {
		return nil, err
	}
	geoj.Coordinates = geom
	return json.Marshal(geoj)
}

func (l LineString) ToGeojson() ([]byte, error) {
	var geoj GeoJson
	geoj.Type = l.Type()
	geom, err := ConvertPoiSet(l)
	if err != nil {
		return nil, err
	}
	geoj.Coordinates = geom
	return json.Marshal(geoj)
}

func (ml MultiLineString) ToGeojson() ([]byte, error) {
	var geoj GeoJson
	geoj.Type = ml.Type()
	geom, err := ConvertPathSet(ml)
	if err != nil {
		return nil, err
	}
	geoj.Coordinates = geom
	return json.Marshal(geoj)
}

func (p Polygon) ToGeojson() ([]byte, error) {
	var geoj GeoJson
	geoj.Type = p.Type()
	geom, err := ConvertPathSet(p)
	if err != nil {
		return nil, err
	}
	geoj.Coordinates = geom
	return json.Marshal(geoj)
}

func (mp MultiPolygon) ToGeojson() ([]byte, error) {
	var geoj GeoJson
	geoj.Type = mp.Type()
	geom, err := ConvertPolygonSet(mp)
	if err != nil {
		return nil, err
	}
	geoj.Coordinates = geom
	return json.Marshal(geoj)
}

// func (c Collection) ToGeojson() ([]byte, error) {
// 	return "GeometryCollection"
// }

func ConvertPoi(p Point) ([]float64, error) {
	return []float64{p.X, p.Y}, nil
}

// only linestring and multipoint
func ConvertPoiSet(g Geometry) ([][]float64, error) {
	var pts []Point
	switch g := g.(type) {
	case LineString:
		pts = g.GetPointSet()
	case MultiPoint:
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

// only multilinestring and polygon
func ConvertPathSet(g Geometry) ([][][]float64, error) {
	var lines []LineString
	switch g := g.(type) {
	case MultiLineString:
		for _, line := range g {
			lines = append(lines, line)
		}
	case Polygon:
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

//only multipolygon
func ConvertPolygonSet(g Geometry) ([][][][]float64, error) {
	var polys []Polygon
	switch g := g.(type) {
	case MultiPolygon:
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

// func decodeGeometries(data interface{}) ([]*Geometry, error) {
// 	if vs, ok := data.([]interface{}); ok {
// 		geometries := make([]*Geometry, 0, len(vs))
// 		for _, v := range vs {
// 			g := &Geometry{}

// 			vmap, ok := v.(map[string]interface{})
// 			if !ok {
// 				break
// 			}

// 			err := decodeGeometry(g, vmap)
// 			if err != nil {
// 				return nil, err
// 			}

// 			geometries = append(geometries, g)
// 		}

// 		if len(geometries) == len(vs) {
// 			return geometries, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("not a valid set of geometries, got %v", data)
// }

// UnmarshalGeometry decodes the data into a GeoJSON geometry.
// Alternately one can call json.Unmarshal(g) directly for the same result.
// func UnmarshalGeometry(data []byte) (*geo.Geometry, error) {
// 	g := &geo.Geometry{}
// 	err := json.Unmarshal(data, g)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return g, nil
// }

// // UnmarshalJSON decodes the data into a GeoJSON geometry.
// // This fulfills the json.Unmarshaler interface.
// func UnmarshalJSON(g *Geometry, data []byte) error {
// 	var object map[string]interface{}
// 	err := json.Unmarshal(data, &object)
// 	if err != nil {
// 		return err
// 	}

// 	return decodeGeometry(g, object)
// }

// // Scan implements the sql.Scanner interface allowing
// // geometry structs to be passed into rows.Scan(...interface{})
// // The columns must be received as GeoJSON Geometry.
// // When using PostGIS a spatial column would need to be wrapped in ST_AsGeoJSON.
// func (g *Geometry) Scan(value interface{}) error {
// 	var data []byte

// 	switch value.(type) {
// 	case string:
// 		data = []byte(value.(string))
// 	case []byte:
// 		data = value.([]byte)
// 	default:
// 		return errors.New("unable to parse this type into geojson")
// 	}

// 	return g.UnmarshalJSON(data)
// }

// func decodeGeometry(g *Geometry, object map[string]interface{}) error {
// 	t, ok := object["type"]
// 	if !ok {
// 		return errors.New("type property not defined")
// 	}

// 	if s, ok := t.(string); ok {
// 		g.Type = GeometryType(s)
// 	} else {
// 		return errors.New("type property not string")
// 	}

// 	bb, err := decodeBoundingBox(object["bbox"])
// 	if err != nil {
// 		return err
// 	}
// 	g.BoundingBox = bb

// 	switch g.Type {
// 	case GeometryPoint:
// 		g.Point, err = decodePosition(object["coordinates"])
// 	case GeometryMultiPoint:
// 		g.MultiPoint, err = decodePositionSet(object["coordinates"])
// 	case GeometryLineString:
// 		g.LineString, err = decodePositionSet(object["coordinates"])
// 	case GeometryMultiLineString:
// 		g.MultiLineString, err = decodePathSet(object["coordinates"])
// 	case GeometryPolygon:
// 		g.Polygon, err = decodePathSet(object["coordinates"])
// 	case GeometryMultiPolygon:
// 		g.MultiPolygon, err = decodePolygonSet(object["coordinates"])
// 	case GeometryCollection:
// 		g.Geometries, err = decodeGeometries(object["geometries"])
// 	}

// 	return err
// }

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

// func decodeGeometries(data interface{}) ([]*Geometry, error) {
// 	if vs, ok := data.([]interface{}); ok {
// 		geometries := make([]*Geometry, 0, len(vs))
// 		for _, v := range vs {
// 			g := &Geometry{}

// 			vmap, ok := v.(map[string]interface{})
// 			if !ok {
// 				break
// 			}

// 			err := decodeGeometry(g, vmap)
// 			if err != nil {
// 				return nil, err
// 			}

// 			geometries = append(geometries, g)
// 		}

// 		if len(geometries) == len(vs) {
// 			return geometries, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("not a valid set of geometries, got %v", data)
// }
