package tests

import (
	"testing"

	gogeo "github.com/sadnessly/go-geo"
)

func TestConvertStringToGeo(t *testing.T) {
	geo, err := gogeo.ConvertStringToGeo("[[31.531234,120.335439],[31.53686,120.344768],[31.560835,120.330253],[31.56173,120.32896],[31.564232,120.318883],[31.564189,120.317951],[31.563801,120.312261],[31.558314,120.296748],[31.558094,120.296301],[31.554538,120.297303],[31.540824,120.307415],[31.535424,120.315807]]",
		gogeo.STR_POIJSON, gogeo.ELEM_POLYGON)
	if err != nil {
		t.Log("error in convert")
	}
	t.Log(gogeo.GeoToWkt(geo))
}
