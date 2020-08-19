package tests

// func TestInpolygon(t *testing.T) {
// 	line := geo.NewLine(geo.Point{2, 1}, geo.Point{8, 1}, geo.Point{8, 6}, geo.Point{6, 6}, geo.Point{6, 3}, geo.Point{4, 3}, geo.Point{4, 5}, geo.Point{2, 5})
// 	poly := *geo.NewPolygon(*geo.NewLinearRingFromLineString(*line))
// 	// verts := [][]float64{{2, 1}, {8, 1}, {8, 6}, {6, 6}, {6, 3}, {4, 3}, {4, 5}, {2, 5}}
// 	t.Log(geo.IsPointInPolygon(geo.Point{5, 4}, poly))
// 	t.Log(geo.IsPointInPolygon(geo.Point{6, 6}, poly))
// 	t.Log(geo.IsPointInPolygon(geo.Point{6, 5}, poly))
// 	t.Log(geo.IsPointInPolygon(geo.Point{7, 4}, poly))
// 	t.Log(geo.IsPointInPolygon(geo.Point{8, 4}, poly))
// 	t.Log(geo.IsPointInPolygon(geo.Point{3, 3}, poly))
// 	t.Log(geo.IsPointInPolygon(geo.Point{4, 3}, poly))

// }

// func TestPolyIntersect(t *testing.T) {
// 	poly1 := *geo.NewPolygonFromPois(geo.Point{32.134057, 118.867052}, geo.Point{32.133586, 118.863761}, geo.Point{32.133369, 118.861667},
// 		geo.Point{32.133224, 118.859658}, geo.Point{32.132899, 118.85782}, geo.Point{32.132899, 118.857221}, geo.Point{32.134257, 118.857245},
// 		geo.Point{32.137153, 118.857095}, geo.Point{32.141044, 118.85701}, //geo.Point{32.141259, 118.859529}, geo.Point{32.141295, 118.86312},
// 		geo.Point{32.14144, 118.866667}, geo.Point{32.1422, 118.86671}, geo.Point{32.143756, 118.866411}, geo.Point{32.145458, 118.866625},
// 		geo.Point{32.145675, 118.867522}, geo.Point{32.145602, 118.872352}, geo.Point{32.145638, 118.875814}, geo.Point{32.145566, 118.884576},
// 		geo.Point{32.145602, 118.888551}, geo.Point{32.147195, 118.895817}, geo.Point{32.145964, 118.895731}, geo.Point{32.140029, 118.894065},
// 		geo.Point{32.136699, 118.892996}, geo.Point{32.134563, 118.892611}, geo.Point{32.133441, 118.893466}, geo.Point{32.128193, 118.898424},
// 		geo.Point{32.126601, 118.899535}, geo.Point{32.120448, 118.903425}, geo.Point{32.116973, 118.90569}, geo.Point{32.116286, 118.905135},
// 		geo.Point{32.112341, 118.900689}, geo.Point{32.109482, 118.898168}, geo.Point{32.106623, 118.89586}, geo.Point{32.104379, 118.894321},
// 		geo.Point{32.10275, 118.892911}, geo.Point{32.101483, 118.891927}, geo.Point{32.100361, 118.890816}, geo.Point{32.098805, 118.888765},
// 		geo.Point{32.097936, 118.886884}, geo.Point{32.09638, 118.881285}, geo.Point{32.095801, 118.878592}, geo.Point{32.095837, 118.876712},
// 		geo.Point{32.096452, 118.875173}, geo.Point{32.096959, 118.874831}, geo.Point{32.098262, 118.87419}, geo.Point{32.102243, 118.872737},
// 		geo.Point{32.107527, 118.870386}, geo.Point{32.114404, 118.866625}, geo.Point{32.116829, 118.865342}, geo.Point{32.117601, 118.864768},
// 		geo.Point{32.119037, 118.863291}, geo.Point{32.122869, 118.860343}, geo.Point{32.125371, 118.85876}, geo.Point{32.126999, 118.858461},
// 		geo.Point{32.128556, 118.858888}, geo.Point{32.129279, 118.860555}, geo.Point{32.130293, 118.863462}, geo.Point{32.131415, 118.866881},
// 		geo.Point{32.131921, 118.867479}, geo.Point{32.134093, 118.867351}, geo.Point{32.134057, 118.867052})

// 	poly2 := *geo.NewPolygonFromPois(geo.Point{32.117047, 118.817769}, geo.Point{32.125698, 118.820974}, geo.Point{32.129955, 118.822375},
// 		geo.Point{32.131159, 118.822599}, geo.Point{32.133231, 118.822706}, geo.Point{32.134679, 118.822471}, geo.Point{32.135512, 118.821915},
// 		geo.Point{32.136598, 118.82121}, geo.Point{32.137653, 118.824106}, geo.Point{32.138521, 118.826372}, geo.Point{32.139236, 118.82868},
// 		geo.Point{32.139508, 118.830593}, geo.Point{32.139779, 118.832142}, geo.Point{32.140326, 118.834077}, geo.Point{32.141014, 118.836471},
// 		geo.Point{32.141833, 118.837581}, geo.Point{32.145172, 118.841952}, geo.Point{32.146665, 118.843545}, geo.Point{32.148592, 118.845212},
// 		geo.Point{32.149569, 118.846099}, geo.Point{32.150811, 118.847029}, geo.Point{32.154918, 118.850448}, geo.Point{32.156601, 118.852927},
// 		geo.Point{32.15756, 118.855129}, geo.Point{32.157995, 118.855963}, geo.Point{32.156203, 118.854915}, geo.Point{32.155244, 118.854317},
// 		geo.Point{32.152439, 118.853697}, geo.Point{32.1494, 118.852949}, geo.Point{32.145744, 118.851987}, geo.Point{32.145654, 118.854188},
// 		geo.Point{32.145694, 118.858699}, geo.Point{32.145599, 118.866585}, geo.Point{32.14426, 118.866392}, geo.Point{32.143573, 118.866285},
// 		geo.Point{32.141456, 118.866456}, geo.Point{32.141125, 118.856838}, geo.Point{32.139126, 118.856945}, geo.Point{32.13699, 118.856988},
// 		geo.Point{32.132755, 118.857116}, geo.Point{32.131886, 118.867332}, geo.Point{32.128629, 118.85874},
// 		geo.Point{32.126493, 118.858142}, geo.Point{32.124684, 118.858933}, geo.Point{32.120395, 118.861968}, geo.Point{32.11647, 118.865217},
// 		geo.Point{32.114008, 118.866499}, geo.Point{32.11017, 118.868593}, geo.Point{32.107165, 118.870261}, geo.Point{32.094847, 118.875561},
// 		geo.Point{32.093779, 118.868679}, geo.Point{32.093642, 118.864362}, geo.Point{32.092555, 118.856261}, geo.Point{32.092139, 118.85173},
// 		geo.Point{32.091505, 118.845404}, geo.Point{32.090727, 118.840317}, geo.Point{32.089317, 118.834389}, geo.Point{32.08909, 118.833671},
// 		geo.Point{32.08888, 118.833008}, geo.Point{32.089748, 118.832767}, geo.Point{32.090046, 118.834023}, geo.Point{32.090481, 118.833718},
// 		geo.Point{32.090916, 118.832671}, geo.Point{32.091524, 118.83194}, geo.Point{32.09348, 118.831576}, geo.Point{32.096657, 118.832046},
// 		geo.Point{32.099246, 118.832815}, geo.Point{32.101944, 118.833756}, geo.Point{32.103683, 118.834119}, geo.Point{32.108679, 118.834226},
// 		geo.Point{32.113604, 118.834418}, geo.Point{32.119642, 118.834568}, geo.Point{32.119425, 118.83335}, geo.Point{32.118393, 118.829973},
// 		geo.Point{32.117216, 118.825421}, geo.Point{32.116872, 118.818688}, geo.Point{32.117047, 118.817769})

// 	// GEO_UNKNOWN = 0
// 	// GEO_DISJOINT = 1
// 	// GEO_CONTAIN = 2
// 	// GEO_EQUAL = 3
// 	// GEO_TOUCH = 4
// 	// GEO_COVER = 5
// 	// GEO_INTERSECT = 6
// 	t.Log(geo.GeoToWkt(poly1))
// 	t.Log(geo.GeoToWkt(poly2))
// 	t.Log("the relation is", geo.PolyRelation(poly1, poly2))
// }

// // func TestSegmentRelation(t *testing.T) {
// // 	relation := geo.SegmentRelation(geo.LineSegment{geo.Point{32.13404, 118.867311}, geo.Point{32.131886, 118.867331}}, geo.LineSegment{geo.Point{32.134057, 118.867052}, geo.Point{32.133586, 118.863761}})
// // 	t.Log(geo.GeoToWkt(geo.LineString{geo.Point{32.13404, 118.867311}, geo.Point{32.131886, 118.867331}}))

// // 	t.Log(geo.GeoToWkt(geo.LineString{geo.Point{32.134057, 118.867052}, geo.Point{32.133586, 118.863761}}))
// // 	t.Log("the relation is", relation)

// // 	rela := geo.Intersect(32.13404, 118.867311, 32.131886, 118.867331, 32.134057, 118.867052, 32.133586, 118.863761)
// // 	t.Log("the relation is", rela)
// // }
