package geo

import (
	"errors"

	"math"
)

//计算三个点的角度 使用公式θ=atan2(v2.y,v2.x)−atan2(v1.y,v1.x)
func CacAngle(point1, centerPoint, point2 Point) (float64, error) {
	v := ConvertPointToVector(point1, centerPoint)
	u := ConvertPointToVector(point2, centerPoint)
	// v := LineToVector(point1, centerPoint)
	// u := LineToVector(point2, centerPoint)
	nrmProduct := v.Length() * u.Length()
	if nrmProduct == 0 {
		err := errors.New("some points is repeat")
		return 0, err
	}
	var theta float64

	theta = math.Atan2(u.Y, u.X) - math.Atan2(v.Y, v.X)
	if theta < 0 {
		theta = theta + 2*math.Pi
	}
	return theta, nil

	// err := errors.New("not validate vector")
	// return 0, err
}

func CacQuadrantAngle(point1, point2 Point) (float64, error) {
	angle, err := CacAngle(Point{point1.X + 0.0000001, point1.Y}, point1, point2)
	return angle, err
}

//求多边形的面积 论文:《多边形面积的计算与面积法的应用》
func GetArea(geo Geometry) float64 {
	switch geo := geo.(type) {
	case Polygon:
		return polyArea(geo)
	case MultiPolygon:
		return MultiPolyArea(geo)
	}
	return 0
}

func polyArea(poly Polygon) float64 {
	lr := poly.GetExteriorRing()
	if lr == nil {
		return 0
	}
	ptCount := lr.GetPointCount() - 1
	var area float64
	for i := 0; i < ptCount; i++ {
		//最后一个点的处理
		j := (i + 1) % ptCount
		area += lr[i].X * lr[j].Y
		area -= lr[i].Y * lr[j].X
	}
	area /= 2
	return math.Abs(area)
}

func MultiPolyArea(multiPoly MultiPolygon) float64 {
	areaSum := 0.0
	for _, v := range multiPoly {
		areaSum += polyArea(v)
	}
	return areaSum
}

//计算多个点的中心
func pointsCenteriod(points ...Point) Point {
	var pointList []Point
	for _, v := range points {
		pointList = append(pointList, v)
	}
	amount := len(pointList)
	if amount == 0 {
		return Point{0, 0}
	}
	var lats, Lat float64
	var lngs, Lng float64
	for _, poi := range pointList {
		lats += poi.X
		lngs += poi.Y
	}
	Lat = lats / float64(amount)
	Lng = lngs / float64(amount)
	return Point{Lat, Lng}
}

//计算顶点的凹凸性 先计算待处理点与相邻点的两个向量，再计算两向量的叉乘，根据求得结果的正负可以判断凹凸性。 0代表凸顶点，1代表凹顶点，2代表平角
func CacConvex(p1, p2, p3 Point) (int8, error) {
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

//Euclidean distance
func PointDistance(p1 Point, p2 Point, srid SRID) float64 {
	if srid == SRID_WGS84_GPS {
		return CoordDistance(p1, p2)
	} else {
		return EuclideanDistance(p1, p2)
	}
}

func EuclideanDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))
}

// //计算点到直线的距离 向量的方法，先求三角形的面积，再用面积除以底边长
func PointToLineDistance(point, p1, p2 Point, srid SRID) float64 {
	if p1.Equal(p2) {
		return PointDistance(p1, point, srid)
	}
	area := polyArea(*NewPolygon(*NewLinearRing(p1, p2, point)))
	dis := PointDistance(p1, p2, srid)
	return 2 * area / dis
}

//计算点到线段的距离 计算了最近点
func PointToSegmentDistance(point, p1, p2 Point, srid SRID) (float64, Point) {
	var xDelta float64 = p2.X - p1.X
	var yDelta float64 = p2.Y - p1.Y

	var u float64 = ((point.X-p1.X)*xDelta + (point.Y-p1.Y)*yDelta) / (xDelta*xDelta + yDelta*yDelta)

	var closestPointOnLine Point
	if u < 0 {
		closestPointOnLine = p1
	} else if u > 1 {
		closestPointOnLine = p2
	} else {
		closestPointOnLine = Point{X: (p1.X + u*xDelta), Y: (p1.Y + u*yDelta)}
	}
	return PointDistance(point, closestPointOnLine, srid), closestPointOnLine
}

//如果是一堆点 即计算其坐标的平均值
func Centroid(geo Geometry) Point {
	switch geo := geo.(type) {
	case Point:
		return geo
	case MultiPoint:
		return pointsCenteriod(geo...)
	case LineString:
		return lineCentroid(geo)
	case Polygon:
		return polyCentroid(geo)
	}
	return Point{0, 0}
}

//TODO:如果线在一条直线上，需要改进算法
func lineCentroid(line LineString) Point {
	return LinearCentroid(*NewRingFromLine(line))
}

func LinearCentroid(ring LinearRing) Point {
	return polyCentroid(*NewPolygon(ring))
}

//https://en.wikipedia.org/wiki/Centroid#Centroid_of_polygon
func polyCentroid(poly Polygon) Point {
	lr := poly.GetExteriorRing()
	if lr == nil {
		return Point{0, 0}
	}
	ptCount := lr.GetPointCount() - 1
	var centroidX, centroidY, signArea float64
	for i := 0; i < ptCount; i++ {
		//最后一个点的处理
		j := (i + 1) % ptCount
		centroidX += (lr[i].X + lr[j].X) * (lr[i].X*lr[j].Y - lr[j].X*lr[i].Y)
		centroidY += (lr[i].Y + lr[j].Y) * (lr[i].X*lr[j].Y - lr[j].X*lr[i].Y)
		signArea += lr[i].X*lr[j].Y - lr[j].X*lr[i].Y
	}
	centroidX *= 1 / (6 * signArea / 2)
	centroidY *= 1 / (6 * signArea / 2)
	return Point{centroidX, centroidY}
}

func PointPolygonDistance(p Point, poly Polygon, srid SRID) float64 {
	if IsPointInPolygon(p, poly) == RELA_CONTAIN || IsPointInPolygon(p, poly) == RELA_TOUCH {
		return 0
	}
	var distance = INF
	lr := poly.GetExteriorRing()
	ptCount := lr.GetPointCount() - 1
	for i := 0; i < ptCount; i++ {
		//最后一个点的处理
		j := (i + 1) % ptCount
		var previousPoint Point = lr[i]
		var currentPoint Point = lr[j]

		segmentDistance, _ := PointToSegmentDistance(p, previousPoint, currentPoint, srid)

		if segmentDistance < distance {
			distance = segmentDistance
		}
	}
	return distance
}

// 由点到线上的最近点
// 返回参数：最近点，点序，距离
func PointHitLineString(p Point, line LineString, srid SRID) (Point, int, float64) {
	distance, nearestIndex, nearestPt := INF, 0, Point{0, 0}
	ptCount := line.GetPointCount()
	for i := 0; i < ptCount-1; i++ {
		var previousPoint Point = line[i]
		var currentPoint Point = line[i+1]

		segmentDistance, tempPt := PointToSegmentDistance(p, previousPoint, currentPoint, srid)

		if segmentDistance < distance {
			distance = segmentDistance
			nearestIndex = i
			nearestPt = tempPt
		}
	}
	return nearestPt, nearestIndex, distance
}
