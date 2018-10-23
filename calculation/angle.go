package calculation

import (
	"errors"

	"math"

	"github.com/sadnessly/go-geo/convert"
	"github.com/sadnessly/go-geo/element"
	"gonum.org/v1/gonum/blas/blas64"
)

//计算三个点的角度 使用公式θ=atan2(v2.y,v2.x)−atan2(v1.y,v1.x)
func CacAngle(point1, centerPoint, point2 element.Point) (float64, error) {
	v := convert.LineToVector(point1, centerPoint)
	u := convert.LineToVector(point2, centerPoint)
	nrmProduct := blas64.Nrm2(2, v) * blas64.Nrm2(2, u)
	if nrmProduct == 0 {
		err := errors.New("some points is repeat")
		return 0, err
	}
	var theta float64
	if u.Inc == 1 && v.Inc == 1 {
		if len(u.Data) == 2 && len(v.Data) == 2 {
			theta = math.Atan2(u.Data[1], u.Data[0]) - math.Atan2(v.Data[1], v.Data[0])
			if theta < 0 {
				theta = theta + 2*math.Pi
			}
			return theta, nil
		} else {
			err := errors.New("no support more than 2 dimensions")
			return 0, err
		}
	}
	err := errors.New("not validate vector")
	return 0, err
}

func CacQuadrantAngle(point1, point2 element.Point) (float64, error) {
	angle, err := CacAngle(element.Point{point1.X + 0.0000001, point1.Y}, point1, point2)
	return angle, err
}
