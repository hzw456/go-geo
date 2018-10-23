package tests

import (
	"testing"

	"github.com/sadnessly/go-geo/element"
	"github.com/sadnessly/go-geo/relation"
)

func TestInpolygon(t *testing.T) {
	line := element.NewLine(element.Point{2, 1}, element.Point{8, 1}, element.Point{8, 6}, element.Point{6, 6}, element.Point{6, 3}, element.Point{4, 3}, element.Point{4, 5}, element.Point{2, 5})
	poly := *element.NewPolygon(*element.NewLinearRing(*line))
	// verts := [][]float64{{2, 1}, {8, 1}, {8, 6}, {6, 6}, {6, 3}, {4, 3}, {4, 5}, {2, 5}}
	t.Log(relation.IsPointInPolygon(element.Point{5, 4}, poly))
	t.Log(relation.IsPointInPolygon(element.Point{6, 6}, poly))
	t.Log(relation.IsPointInPolygon(element.Point{6, 5}, poly))
	t.Log(relation.IsPointInPolygon(element.Point{7, 4}, poly))
	t.Log(relation.IsPointInPolygon(element.Point{8, 4}, poly))
	t.Log(relation.IsPointInPolygon(element.Point{3, 3}, poly))
	t.Log(relation.IsPointInPolygon(element.Point{4, 3}, poly))

}
