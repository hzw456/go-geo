package tests

import (
	"testing"

	"git.xiaojukeji.com/haozhiwei/go-geo/calculation"
	"git.xiaojukeji.com/haozhiwei/go-geo/convert"
	"git.xiaojukeji.com/haozhiwei/go-geo/element"
)

func TestDistance(t *testing.T) {
	newPoint := element.NewPoint(153.101112401, 27.797998206)
	newPoint1 := element.NewPoint(200, 200)
	//newPoint.SetX(500)
	t.Log(calculation.PointDistance(*newPoint, *newPoint1))
	t.Log(calculation.PointsCenteriod(*newPoint, *newPoint1))
	poly := newPoint.Buffer(0.5)
	t.Log(poly)
	str, _ := convert.PolygonToWkt(poly)
	t.Logf(str)
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
	poly := element.NewPolygon(*lr)
	t.Log(poly.Area())
}

func test