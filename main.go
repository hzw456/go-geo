package main

import (
	"fmt"
	"math"

	"github.com/go-geo/buffer"
	"github.com/go-geo/convert"
	"github.com/go-geo/element"
	"gonum.org/v1/gonum/blas/blas64"
)

func main() {
	v := blas64.Vector{Inc: 1, Data: []float64{1, 1}}
	u := blas64.Vector{[]float64{2, 0}, 1}
	//blas64.Dot(1, v, u)
	//theta1 = acosd(dot([x1-x2,y1-y2],[x3-x2,y3-y2])/(norm([x1-x2,y1-y2])*norm([x1-x2,y3-y2])));
	theta1 := math.Acos(blas64.Dot(2, v, u)/(blas64.Nrm2(2, v)*blas64.Nrm2(2, u))) * 180 / math.Pi
	fmt.Println("v has length:", blas64.Nrm2(2, v), blas64.Nrm2(2, u), blas64.Dot(2, v, u), theta1)
	newLine := element.NewLine(element.Point{0, 0}, element.Point{5, 2}, element.Point{7, 2})
	fmt.Println(convert.PolygonToWkt(buffer.Buffer(*newLine, 0.5)))
}
