package geojson

import (
	"testing"

	geo "github.com/hzw456/go-geo"
)

func TestUnmarshalGeoPoint(t *testing.T) {
	rawJSON := `
	  { "type": "Feature",
	    "id": 123,
	    "geometry": {"type": "Point", "coordinates": [102.0, 0.5]}
	  }`

	f, err := UnmarshalFeature([]byte(rawJSON))
	if err != nil {
		t.Fatalf("should unmarshal feature without issue, err %v", err)
	}
	geom, err := UnmarshalGeo(f.GeometryJson)
	if err != nil {
		t.Fatalf("should unmarshal Geo without issue, err %v", err)
	}
	gpoint := geom.(geo.Point)
	if gpoint.X != 102.0 || gpoint.Y != 0.5 {
		t.Fatalf("should Geo Point Has issue %v", gpoint)
	}
}

func TestUnmarshalGeoPointZ(t *testing.T) {
	rawJSON := `
	  { "type": "Feature",
	    "id": 123,
	    "geometry": {"type": "Point", "coordinates": [102.0, 0.5, 12.0]}
	  }`

	f, err := UnmarshalFeature([]byte(rawJSON))
	if err != nil {
		t.Fatalf("should unmarshal feature without issue, err %v", err)
	}
	geom, err := UnmarshalGeo(f.GeometryJson)
	if err != nil {
		t.Fatalf("should unmarshal Geo without issue, err %v", err)
	}
	gpoint := geom.(geo.PointZ)
	if gpoint.X != 102.0 || gpoint.Y != 0.5 || gpoint.Z != 12.0 {
		t.Fatalf("should Geo Point Has issue %v", gpoint)
	}
}
