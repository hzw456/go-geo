package gogeo

import (
	"encoding/json"
)

// string解析成geo 目标支持 geojson，wkt,wkb等标准的gis类型（GeometryType可以传UNKNOWN,以string中的类型为准）
// 普通的poi json串（需要提供要素类型）
// TODO :目前只准备支持面和json格式的poilist 其他的待后续支持
func ConvertStringToGeo(value string, strType GeoStringType, geoType GeometryType) (Geometry, error) {
	if strType == STR_POIJSON && geoType == ELEM_POLYGON {
		var result [][]float64
		if err := json.Unmarshal([]byte(value), &result); err != nil {
			return nil, err
		}
		var pois []Point
		for i, _ := range result {
			pois = append(pois, Point{result[i][0], result[i][1]})
		}
		return *NewPolygonFromPois(pois...), nil
	}

	return nil, nil
}
