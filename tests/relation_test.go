package tests

import (
	"testing"

	gogeo "github.com/sadnessly/go-geo"
)

func TestInpolygon(t *testing.T) {
	line := gogeo.NewLine(gogeo.Point{2, 1}, gogeo.Point{8, 1}, gogeo.Point{8, 6}, gogeo.Point{6, 6}, gogeo.Point{6, 3}, gogeo.Point{4, 3}, gogeo.Point{4, 5}, gogeo.Point{2, 5})
	poly := *gogeo.NewPolygon(*gogeo.NewLinearRing(*line))
	// verts := [][]float64{{2, 1}, {8, 1}, {8, 6}, {6, 6}, {6, 3}, {4, 3}, {4, 5}, {2, 5}}
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{5, 4}, poly))
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{6, 6}, poly))
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{6, 5}, poly))
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{7, 4}, poly))
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{8, 4}, poly))
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{3, 3}, poly))
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{4, 3}, poly))

}
