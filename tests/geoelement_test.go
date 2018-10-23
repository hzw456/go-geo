package tests

import (
	"fmt"
	"testing"

	"github.com/sadnessly/go-geo/buffer"
	"github.com/sadnessly/go-geo/calculation"
	"github.com/sadnessly/go-geo/convert"
	"github.com/sadnessly/go-geo/element"
	"github.com/sadnessly/go-geo/simplify"
)

func TestDistance(t *testing.T) {
	newPoint := element.NewPoint(153.101112401, 27.797998206)
	newPoint1 := element.NewPoint(200, 200)
	//newPoint.SetX(500)
	t.Log(calculation.PointDistance(*newPoint, *newPoint1))
	t.Log(calculation.PointsCenteriod(*newPoint, *newPoint1))
}

func TestLength(t *testing.T) {
	newPoint := element.NewPoint(153.101112401, 27.797998206)
	newPoint1 := element.NewPoint(200, 200)
	line := element.NewLine(*newPoint, *newPoint1)
	t.Log(line.Length())
}

func TestArea(t *testing.T) {
	newPoint1 := element.NewPoint(100, 100)
	newPoint2 := element.NewPoint(200, 100)
	newPoint3 := element.NewPoint(200, 200)
	newPoint4 := element.NewPoint(100, 200)
	lr := element.NewLinearRing(*element.NewLine(*newPoint1, *newPoint2, *newPoint3, *newPoint4))
	poly := *element.NewPolygon(*lr)
	t.Log(calculation.Area(poly))
}

func TestBuffer(t *testing.T) {
	newPoint1 := element.NewPoint(200, 200)
	buffer.Buffer(*newPoint1, 10)
	t.Log(convert.PolygonToWkt(buffer.Buffer(*newPoint1, 10)))
	newLine := element.NewLine(element.Point{0, 0}, element.Point{5, 2}, element.Point{7, 2}, element.Point{10, 12})
	t.Log(convert.PolygonToWkt(buffer.Buffer(*newLine, 5)))
}

func TestSimplify(t *testing.T) {
	newPoint1 := element.NewPoint(0, 0)
	newPoint2 := element.NewPoint(0.5, 0.5)
	newPoint3 := element.NewPoint(1, 1)
	newPoint4 := element.NewPoint(2, 2)
	newPoint5 := element.NewPoint(2, 5)
	line := *element.NewLine(*newPoint1, *newPoint2, *newPoint3, *newPoint4, *newPoint5)
	fmt.Println(simplify.DouglasPeucker(1).Simplify(line))
}
