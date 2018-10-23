package calculation

import (
	"math"

	"github.com/sadnessly/go-geo/element"
)

//计算顶点的凹凸性 先计算待处理点与相邻点的两个向量，再计算两向量的叉乘，根据求得结果的正负可以判断凹凸性。 0代表凸顶点，1代表凹顶点，2代表平角
func CacConvex(p1, p2, p3 element.Point) (int8, error) {
	//直接采用算sin theata 来判断凹凸性
	theata, err := CacAngle(p1, p2, p3)
	if err != nil {
		return -1, err
	}
	res := math.Sin(theata)
	if res < 0 {
		return 0, nil
	} else if res > 0 {
		return 1, nil
	}
	return 2, nil
}
