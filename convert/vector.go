package convert

import (
	"github.com/sadnessly/go-geo/element"
	"gonum.org/v1/gonum/blas/blas64"
)

//向量转换相关的类

//点转换为向量 两个点所组成的向量，方向为第一个点指向第二个点
func LineToVector(linePoint1, linePoint2 element.Point) blas64.Vector {
	return blas64.Vector{Data: []float64{linePoint2.X - linePoint1.X, linePoint2.Y - linePoint1.Y}, Inc: 1}
}
