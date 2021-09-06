package geojson

import (
	"encoding/json"

	geo "github.com/hzw456/go-geo"
)

type Feature struct {
	ID           interface{}            `json:"id,omitempty"`
	Type         string                 `json:"type"`
	BoundingBox  []float64              `json:"bbox,omitempty"`
	GeometryJson *GeoJson               `json:"geometry"`
	Properties   map[string]interface{} `json:"properties"`
	CRS          map[string]interface{} `json:"crs,omitempty"`
	Geometry     geo.Geometry           `json:"-"`
}

// NewFeature creates and initializes a GeoJSON feature given the required attributes.
func NewFeature(geometry geo.Geometry) *Feature {
	gjson, _ := MarshalGeo(geometry)
	var geoj GeoJson
	json.Unmarshal(gjson, &geoj)
	return &Feature{
		Type:         "Feature",
		GeometryJson: &geoj,
		Properties:   make(map[string]interface{}),
		Geometry:     geometry,
	}
}

func (f Feature) SetBoundingBox() {
	box := geo.BoundingBox(f.Geometry)
	f.BoundingBox = []float64{box.MinX, box.MinY, box.MaxX, box.MaxY}
}

// MarshalJSON converts the feature object into the proper JSON.
// It will handle the encoding of all the child geometries.
// Alternately one can call json.Marshal(f) directly for the same result.
func (f Feature) MarshalJSON() ([]byte, error) {
	type feature Feature
	fea := &feature{
		ID:           f.ID,
		Type:         "Feature",
		GeometryJson: f.GeometryJson,
	}

	if f.BoundingBox != nil && len(f.BoundingBox) != 0 {
		fea.BoundingBox = f.BoundingBox
	}
	if f.Properties != nil && len(f.Properties) != 0 {
		fea.Properties = f.Properties
	}
	if f.CRS != nil && len(f.CRS) != 0 {
		fea.CRS = f.CRS
	}

	return json.Marshal(fea)
}

// UnmarshalFeature decodes the data into a GeoJSON feature.
// Alternately one can call json.Unmarshal(f) directly for the same result.
func UnmarshalFeature(data []byte) (*Feature, error) {
	f := &Feature{}
	err := json.Unmarshal(data, f)
	if err != nil {
		return nil, err
	}

	return f, nil
}
