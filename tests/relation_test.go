package tests

import (
	"testing"

	gogeo "github.com/sadnessly/go-geo"
)

func TestInpolygon(t *testing.T) {
	line := gogeo.NewLine(gogeo.Point{2, 1}, gogeo.Point{8, 1}, gogeo.Point{8, 6}, gogeo.Point{6, 6}, gogeo.Point{6, 3}, gogeo.Point{4, 3}, gogeo.Point{4, 5}, gogeo.Point{2, 5})
	poly := *gogeo.NewPolygon(*gogeo.NewLinearRingFromLineString(*line))
	// verts := [][]float64{{2, 1}, {8, 1}, {8, 6}, {6, 6}, {6, 3}, {4, 3}, {4, 5}, {2, 5}}
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{5, 4}, poly))
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{6, 6}, poly))
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{6, 5}, poly))
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{7, 4}, poly))
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{8, 4}, poly))
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{3, 3}, poly))
	t.Log(gogeo.IsPointInPolygon(gogeo.Point{4, 3}, poly))

}

func TestPolyIntersect(t *testing.T) {
	poly1 := *gogeo.NewPolygonFromPois(gogeo.Point{32.134057, 118.867052}, gogeo.Point{32.133586, 118.863761}, gogeo.Point{32.133369, 118.861667},
		gogeo.Point{32.133224, 118.859658}, gogeo.Point{32.132899, 118.85782}, gogeo.Point{32.132899, 118.857221}, gogeo.Point{32.134257, 118.857245},
		gogeo.Point{32.137153, 118.857095}, gogeo.Point{32.141044, 118.85701}, //gogeo.Point{32.141259, 118.859529}, gogeo.Point{32.141295, 118.86312},
		gogeo.Point{32.14144, 118.866667}, gogeo.Point{32.1422, 118.86671}, gogeo.Point{32.143756, 118.866411}, gogeo.Point{32.145458, 118.866625},
		gogeo.Point{32.145675, 118.867522}, gogeo.Point{32.145602, 118.872352}, gogeo.Point{32.145638, 118.875814}, gogeo.Point{32.145566, 118.884576},
		gogeo.Point{32.145602, 118.888551}, gogeo.Point{32.147195, 118.895817}, gogeo.Point{32.145964, 118.895731}, gogeo.Point{32.140029, 118.894065},
		gogeo.Point{32.136699, 118.892996}, gogeo.Point{32.134563, 118.892611}, gogeo.Point{32.133441, 118.893466}, gogeo.Point{32.128193, 118.898424},
		gogeo.Point{32.126601, 118.899535}, gogeo.Point{32.120448, 118.903425}, gogeo.Point{32.116973, 118.90569}, gogeo.Point{32.116286, 118.905135},
		gogeo.Point{32.112341, 118.900689}, gogeo.Point{32.109482, 118.898168}, gogeo.Point{32.106623, 118.89586}, gogeo.Point{32.104379, 118.894321},
		gogeo.Point{32.10275, 118.892911}, gogeo.Point{32.101483, 118.891927}, gogeo.Point{32.100361, 118.890816}, gogeo.Point{32.098805, 118.888765},
		gogeo.Point{32.097936, 118.886884}, gogeo.Point{32.09638, 118.881285}, gogeo.Point{32.095801, 118.878592}, gogeo.Point{32.095837, 118.876712},
		gogeo.Point{32.096452, 118.875173}, gogeo.Point{32.096959, 118.874831}, gogeo.Point{32.098262, 118.87419}, gogeo.Point{32.102243, 118.872737},
		gogeo.Point{32.107527, 118.870386}, gogeo.Point{32.114404, 118.866625}, gogeo.Point{32.116829, 118.865342}, gogeo.Point{32.117601, 118.864768},
		gogeo.Point{32.119037, 118.863291}, gogeo.Point{32.122869, 118.860343}, gogeo.Point{32.125371, 118.85876}, gogeo.Point{32.126999, 118.858461},
		gogeo.Point{32.128556, 118.858888}, gogeo.Point{32.129279, 118.860555}, gogeo.Point{32.130293, 118.863462}, gogeo.Point{32.131415, 118.866881},
		gogeo.Point{32.131921, 118.867479}, gogeo.Point{32.134093, 118.867351}, gogeo.Point{32.134057, 118.867052})

	poly2 := *gogeo.NewPolygonFromPois(gogeo.Point{32.117047, 118.817769}, gogeo.Point{32.125698, 118.820974}, gogeo.Point{32.129955, 118.822375},
		gogeo.Point{32.131159, 118.822599}, gogeo.Point{32.133231, 118.822706}, gogeo.Point{32.134679, 118.822471}, gogeo.Point{32.135512, 118.821915},
		gogeo.Point{32.136598, 118.82121}, gogeo.Point{32.137653, 118.824106}, gogeo.Point{32.138521, 118.826372}, gogeo.Point{32.139236, 118.82868},
		gogeo.Point{32.139508, 118.830593}, gogeo.Point{32.139779, 118.832142}, gogeo.Point{32.140326, 118.834077}, gogeo.Point{32.141014, 118.836471},
		gogeo.Point{32.141833, 118.837581}, gogeo.Point{32.145172, 118.841952}, gogeo.Point{32.146665, 118.843545}, gogeo.Point{32.148592, 118.845212},
		gogeo.Point{32.149569, 118.846099}, gogeo.Point{32.150811, 118.847029}, gogeo.Point{32.154918, 118.850448}, gogeo.Point{32.156601, 118.852927},
		gogeo.Point{32.15756, 118.855129}, gogeo.Point{32.157995, 118.855963}, gogeo.Point{32.156203, 118.854915}, gogeo.Point{32.155244, 118.854317},
		gogeo.Point{32.152439, 118.853697}, gogeo.Point{32.1494, 118.852949}, gogeo.Point{32.145744, 118.851987}, gogeo.Point{32.145654, 118.854188},
		gogeo.Point{32.145694, 118.858699}, gogeo.Point{32.145599, 118.866585}, gogeo.Point{32.14426, 118.866392}, gogeo.Point{32.143573, 118.866285},
		gogeo.Point{32.141456, 118.866456}, gogeo.Point{32.141125, 118.856838}, gogeo.Point{32.139126, 118.856945}, gogeo.Point{32.13699, 118.856988},
		gogeo.Point{32.132755, 118.857116}, gogeo.Point{32.131886, 118.867332}, gogeo.Point{32.128629, 118.85874},
		gogeo.Point{32.126493, 118.858142}, gogeo.Point{32.124684, 118.858933}, gogeo.Point{32.120395, 118.861968}, gogeo.Point{32.11647, 118.865217},
		gogeo.Point{32.114008, 118.866499}, gogeo.Point{32.11017, 118.868593}, gogeo.Point{32.107165, 118.870261}, gogeo.Point{32.094847, 118.875561},
		gogeo.Point{32.093779, 118.868679}, gogeo.Point{32.093642, 118.864362}, gogeo.Point{32.092555, 118.856261}, gogeo.Point{32.092139, 118.85173},
		gogeo.Point{32.091505, 118.845404}, gogeo.Point{32.090727, 118.840317}, gogeo.Point{32.089317, 118.834389}, gogeo.Point{32.08909, 118.833671},
		gogeo.Point{32.08888, 118.833008}, gogeo.Point{32.089748, 118.832767}, gogeo.Point{32.090046, 118.834023}, gogeo.Point{32.090481, 118.833718},
		gogeo.Point{32.090916, 118.832671}, gogeo.Point{32.091524, 118.83194}, gogeo.Point{32.09348, 118.831576}, gogeo.Point{32.096657, 118.832046},
		gogeo.Point{32.099246, 118.832815}, gogeo.Point{32.101944, 118.833756}, gogeo.Point{32.103683, 118.834119}, gogeo.Point{32.108679, 118.834226},
		gogeo.Point{32.113604, 118.834418}, gogeo.Point{32.119642, 118.834568}, gogeo.Point{32.119425, 118.83335}, gogeo.Point{32.118393, 118.829973},
		gogeo.Point{32.117216, 118.825421}, gogeo.Point{32.116872, 118.818688}, gogeo.Point{32.117047, 118.817769})

	// GEO_UNKNOWN = 0
	// GEO_DISJOINT = 1
	// GEO_CONTAIN = 2
	// GEO_EQUAL = 3
	// GEO_TOUCH = 4
	// GEO_COVER = 5
	// GEO_INTERSECT = 6
	t.Log(gogeo.GeoToWkt(poly1))
	t.Log(gogeo.GeoToWkt(poly2))
	t.Log("the relation is", gogeo.PolyRelation(poly1, poly2))
}

// func TestSegmentRelation(t *testing.T) {
// 	relation := gogeo.SegmentRelation(gogeo.LineSegment{gogeo.Point{32.13404, 118.867311}, gogeo.Point{32.131886, 118.867331}}, gogeo.LineSegment{gogeo.Point{32.134057, 118.867052}, gogeo.Point{32.133586, 118.863761}})
// 	t.Log(gogeo.GeoToWkt(gogeo.LineString{gogeo.Point{32.13404, 118.867311}, gogeo.Point{32.131886, 118.867331}}))

// 	t.Log(gogeo.GeoToWkt(gogeo.LineString{gogeo.Point{32.134057, 118.867052}, gogeo.Point{32.133586, 118.863761}}))
// 	t.Log("the relation is", relation)

// 	rela := gogeo.Intersect(32.13404, 118.867311, 32.131886, 118.867331, 32.134057, 118.867052, 32.133586, 118.863761)
// 	t.Log("the relation is", rela)
// }
