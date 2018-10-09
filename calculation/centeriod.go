package calculation

import "go-geo/element"

//计算多个点的中心
func PointsCenteriod(points ...element.Point) *element.Point {
	var pointList []element.Point
	for _, v := range points {
		pointList = append(pointList, v)
	}
	return PointListCenteriod(pointList)
}

func PointListCenteriod(pointList []element.Point) *element.Point {
	amount := len(pointList)
	if amount == 0 {
		return nil
	}
	var lats, Lat float64
	var lngs, Lng float64
	for _, poi := range pointList {
		lats += poi.X
		lngs += poi.Y
	}
	Lat = lats / float64(amount)
	Lng = lngs / float64(amount)
	return &element.Point{Lat, Lng}
}
