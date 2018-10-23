package simplify

import (
	"sort"

	"github.com/sadnessly/go-geo/calculation"
	"github.com/sadnessly/go-geo/element"
)

//道格拉斯-匹克算法对线简化，参数：简化阈值 参考网站：https://blog.csdn.net/deram_boy/article/details/39177015
type DouglasPeuckerSimplifier struct {
	Threshold float64
}

func DouglasPeucker(threshold float64) *DouglasPeuckerSimplifier {
	return &DouglasPeuckerSimplifier{
		Threshold: threshold,
	}
}

func (s *DouglasPeuckerSimplifier) Simplify(line element.LineString) element.LineString {
	///获取需要删除的点的序号
	delIndexs := dpWorker(line, s.Threshold)
	//排序，从后往前删
	sort.Sort(sort.Reverse(sort.IntSlice(delIndexs)))

	for _, v := range delIndexs {
		line = append(line[:v], line[v+1:]...)
	}
	return line
}

func dpWorker(line element.LineString, threshold float64) []int {
	var stack []int
	stack = append(stack, 0, len(line)-1)
	var delIndexs []int
	for len(stack) > 1 {
		start := stack[len(stack)-2]
		end := stack[len(stack)-1]

		// modify the line in place
		maxDist := 0.0
		maxIndex := 0

		for i := start + 1; i < end; i++ {
			dist := calculation.PointToLineDistance(line[i], line[start], line[end])
			if dist > maxDist {
				maxDist = dist
				maxIndex = i
			}
		}
		if maxDist > threshold {
			stack[len(stack)-1] = maxIndex
			stack = append(stack, maxIndex, end)
		} else {
			for i := start + 1; i < end; i++ {
				delIndexs = append(delIndexs, i)
			}
			stack = stack[:len(stack)-2]
		}
	}
	return delIndexs
}
