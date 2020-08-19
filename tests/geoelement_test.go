package tests

// func TestDistance(t *testing.T) {
// 	newPoint := geo.NewPoint(153.101112401, 27.797998206)
// 	newPoint1 := geo.NewPoint(200, 200)
// 	//newPoint.SetX(500)
// 	t.Log(geo.PointDistance(*newPoint, *newPoint1))
// 	//t.Log(geo.Centroid(*newPoint, *newPoint1))
// }

// func TestLength(t *testing.T) {
// 	newPoint := geo.NewPoint(153.101112401, 27.797998206)
// 	newPoint1 := geo.NewPoint(200, 200)
// 	line := geo.NewLine(*newPoint, *newPoint1)
// 	t.Log(line.Length())
// }

// func TestArea(t *testing.T) {
// 	newPoint1 := geo.NewPoint(100, 100)
// 	newPoint2 := geo.NewPoint(200, 100)
// 	newPoint3 := geo.NewPoint(200, 200)
// 	newPoint4 := geo.NewPoint(100, 200)
// 	lr := geo.NewLinearRing(*newPoint1, *newPoint2, *newPoint3, *newPoint4)
// 	poly := *geo.NewPolygon(*lr)
// 	t.Log(geo.GetArea(poly))
// }

// func TestSimplify(t *testing.T) {
// 	newPoint1 := geo.NewPoint(0, 0)
// 	newPoint2 := geo.NewPoint(0.5, 0.5)
// 	newPoint3 := geo.NewPoint(1, 1)
// 	newPoint4 := geo.NewPoint(2, 2)
// 	newPoint5 := geo.NewPoint(3, 3)
// 	line := *geo.NewLine(*newPoint1, *newPoint2, *newPoint3, *newPoint4, *newPoint5)
// 	fmt.Println(geo.DouglasPeuckerSimplifier{1}.Simplify(line))
// }

// func TestPointToLineDis(t *testing.T) {
// 	newPoint1 := *geo.NewPoint(0, 0)
// 	newPoint2 := *geo.NewPoint(1, 2)
// 	newPoint3 := *geo.NewPoint(2, 5)
// 	fmt.Println(geo.PointToLineDistance(newPoint2, newPoint1, newPoint3))
// }

// func TestWkt(t *testing.T) {
// 	newPoint1 := *geo.NewPoint(0, 0)
// 	newPoint2 := *geo.NewPoint(1, 2)
// 	newPoint3 := *geo.NewPoint(2, 2)
// 	newPoint4 := *geo.NewPoint(2, 5)
// 	fmt.Println(geo.GeoToWkt(geo.MultiPoint{newPoint1, newPoint2, newPoint3}))
// 	line1 := *geo.NewLine(newPoint1, newPoint2, newPoint3)
// 	line2 := *geo.NewLine(newPoint2, newPoint3, newPoint4)
// 	poly1 := *geo.NewPolygon(*geo.NewLinearRingFromLineString(line1))
// 	poly2 := *geo.NewPolygon(*geo.NewLinearRingFromLineString(line2))
// 	fmt.Println(geo.GeoToWkt(*geo.NewMultiPolygon(poly1, poly2)))
// }

// func TestEnvolope(t *testing.T) {
// 	newPoint1 := *geo.NewPoint(0, 0)
// 	newPoint2 := *geo.NewPoint(1, 2)
// 	newPoint3 := *geo.NewPoint(2, 2)
// 	//newPoint4 := *geo.NewPoint(2, 5)
// 	line1 := *geo.NewLine(newPoint1, newPoint2, newPoint3)
// 	//line2 := *geo.NewLine(newPoint2, newPoint3, newPoint4)
// 	poly1 := *geo.NewPolygon(*geo.NewLinearRingFromLineString(line1))
// 	//poly2 := *geo.NewPolygon(*geo.NewLinearRing(line2))

// 	fmt.Println(geo.GeoToWkt(geo.BoxToGeo(geo.Envelope(poly1))))
// }

// func TestPointinPoly(t *testing.T) {
// 	newPoint1 := *geo.NewPoint(0, 0)
// 	newPoint2 := *geo.NewPoint(1, 2)
// 	newPoint3 := *geo.NewPoint(2, 2)
// 	newPoint4 := *geo.NewPoint(2, 5)
// 	newPoint5 := *geo.NewPoint(0.5, 0.7)
// 	newPoint6 := *geo.NewPoint(0.5, 0.5)
// 	line1 := *geo.NewLine(newPoint1, newPoint2, newPoint3)
// 	poly1 := *geo.NewPolygon(*geo.NewLinearRingFromLineString(line1))
// 	if geo.IsPointInPolygon(newPoint4, poly1) == geo.RELA_CONTAIN {
// 		t.Error("failed, the point is not in poly")
// 	} else {
// 		t.Log("success")
// 	}

// 	if geo.IsPointInPolygon(newPoint5, poly1) == geo.RELA_DISJOINT {
// 		t.Error("failed, the point is in poly")
// 	} else {
// 		t.Log("success")
// 	}

// 	if geo.IsPointInPolygon(newPoint6, poly1) == geo.RELA_CONTAIN {
// 		t.Error("failed, the point is not in poly")
// 	} else {
// 		t.Log("success")
// 	}
// }
