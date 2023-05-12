package tests

import (
	"testing"

	"github.com/hzw456/go-geo"
)

func TestInpolygon(t *testing.T) {
	line := geo.NewLineString(geo.Point{2, 1}, geo.Point{8, 1}, geo.Point{8, 6}, geo.Point{6, 6}, geo.Point{6, 3}, geo.Point{4, 3}, geo.Point{4, 5}, geo.Point{2, 5})
	poly := *geo.NewPolygon(*geo.NewRingFromLine(*line))
	// verts := [][]float64{{2, 1}, {8, 1}, {8, 6}, {6, 6}, {6, 3}, {4, 3}, {4, 5}, {2, 5}}
	t.Log(geo.IsPointInPolygon(geo.Point{5, 4}, poly))
	t.Log(geo.IsPointInPolygon(geo.Point{6, 6}, poly))
	t.Log(geo.IsPointInPolygon(geo.Point{6, 5}, poly))
	t.Log(geo.IsPointInPolygon(geo.Point{7, 4}, poly))
	t.Log(geo.IsPointInPolygon(geo.Point{8, 4}, poly))
	t.Log(geo.IsPointInPolygon(geo.Point{3, 3}, poly))
	t.Log(geo.IsPointInPolygon(geo.Point{4, 3}, poly))

}
