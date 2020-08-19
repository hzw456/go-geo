package geos

import (
	"encoding/hex"
	"math"
	"strings"
	"testing"
)

func TestGeometryTypeConsts(t *testing.T) {
	if POINT.ToString() != "Point" {
		t.Errorf("Error: GeometryType POINT define error")
	}

	if LINESTRING.ToString() != "LineString" {
		t.Errorf("Error: GeometryType LINESTRING define error")
	}

	if LINEARRING.ToString() != "LinearRing" {
		t.Errorf("Error: GeometryType LINEARRING define error")
	}

	if POLYGON.ToString() != "Polygon" {
		t.Errorf("Error: GeometryType POLYGON define error")
	}

	if MULTIPOINT.ToString() != "MultiPoint" {
		t.Errorf("Error: GeometryType MULTIPOINT define error")
	}

	if MULTILINESTRING.ToString() != "MultiLineString" {
		t.Errorf("Error: GeometryType MULTILINESTRING define error")
	}

	if MULTIPOLYGON.ToString() != "MultiPolygon" {
		t.Errorf("Error: GeometryType MULTIPOLYGON define error")
	}

	if GEOMETRYCOLLECTION.ToString() != "GeometryCollection" {
		t.Errorf("Error: GeometryType GEOMETRYCOLLECTION define error")
	}
}

func TestCreateFromWKT(t *testing.T) {
	wkt := "POINT (0 0)"
	geom := CreateFromWKT(wkt)

	if geom == nil {
		t.Errorf("Error: CreateFromWKT(%q) error", wkt)
	}
}

func TestCreateFromWKB(t *testing.T) {
	pt := CreatePoint(0, 0)
	wkb := pt.ToWKB()
	geom := CreateFromWKB(wkb)

	if geom == nil {
		t.Errorf("Error: CreateFromWKB(%q) error", hex.EncodeToString(wkb))
	}
}

func TestCreatePoint(t *testing.T) {
	pt2d := CreatePoint(116.39, 39.9)
	if pt2d != nil && pt2d.GetType() == POINT {
		t.Logf("Log: CreatePoint(116.39, 39.9) returns %q", pt2d.ToWKT())
	} else {
		t.Errorf("Error: CreatePoint(116.39, 39.9) error")
	}

	pt3d := CreatePointZ(116, 40, 1)
	if pt3d != nil && pt3d.GetType() == POINT {
		t.Logf("Log: CreatePointZ(116, 40, 1) returns %q", pt3d.ToWKT())
	} else {
		t.Errorf("Error: CreatePointZ(116, 40, 1) error")
	}
}

func TestCreateLinearRing(t *testing.T) {
	coords := []Point{
		Point{0, 0},
		Point{1, 0},
		Point{1, 1},
		Point{0, 1},
		Point{0, 0}}

	geom := CreateLinearRing(coords)
	if geom != nil && geom.GetType() == LINEARRING {
		t.Logf("Log: CreateLinearRing([[0,0], [1,0], [1,1], [0,1], [0,0]]) returns %q", geom.ToWKT())
	} else {
		t.Errorf("Error: CreateLinearRing([[0,0], [1,0], [1,1], [0,1], [0,0]]) error")
	}
}

func TestCreateLinearRingZ(t *testing.T) {
	coords := []CoordZ{
		CoordZ{0, 0, 0},
		CoordZ{1, 0, 0},
		CoordZ{1, 1, 0},
		CoordZ{0, 1, 0},
		CoordZ{0, 0, 0}}

	geom := CreateLinearRingZ(coords)
	if geom != nil && geom.GetType() == LINEARRING {
		t.Logf("Log: CreateLinearRingZ([[0,0,0], [1,0,0], [1,1,0], [0,1,0], [0,0,0]]) returns %q", geom.ToWKT())
	} else {
		t.Errorf("Error: CreateLinearRingZ([[0,0,0], [1,0,0], [1,1,0], [0,1,0], [0,0,0]]) error")
	}
}

func TestCreateLineString(t *testing.T) {
	coords := []Point{
		Point{0, 0},
		Point{1, 1},
		Point{2, 1}}

	geom := CreateLineString(coords)
	if geom != nil && geom.GetType() == LINESTRING {
		t.Logf("Log: CreateLineString([[0,0], [1,1], [2,1]]) returns %q", geom.ToWKT())
	} else {
		t.Errorf("Error: CreateLineString([[0,0], [1,1], [2,1]]) error")
	}
}

func TestCreateLineStringZ(t *testing.T) {
	coords := []CoordZ{
		CoordZ{0, 0, 0},
		CoordZ{1, 1, 0},
		CoordZ{2, 1, 0}}

	geom := CreateLineStringZ(coords)
	if geom != nil && geom.GetType() == LINESTRING {
		t.Logf("Log: CreateLineStringZ([[0,0,0], [1,1,0], [2,1,0]]) returns %q", geom.ToWKT())
	} else {
		t.Errorf("Error: CreateLineStringZ([[0,0,0], [1,1,0], [2,1,0]]) error")
	}
}

func TestCreatePolygon(t *testing.T) {
	shell := []Point{
		Point{0, 0},
		Point{3, 0},
		Point{3, 3},
		Point{0, 3},
		Point{0, 0}}

	hole := []Point{
		Point{1, 1},
		Point{2, 1},
		Point{2, 2},
		Point{1, 2},
		Point{1, 1}}

	geom := CreatePolygon(shell, hole)
	if geom != nil && geom.GetType() == POLYGON {
		t.Logf("Log: CreatePolygon([[0,0], [3,0], [3,3], [0,3], [0,0]], [[1,1], [2,1], [2,2], [1,2], [1,1]]) returns %q", geom.ToWKT())
	} else {
		t.Errorf("Error: CreatePolygon([[0,0], [3,0], [3,3], [0,3], [0,0]], [[1,1], [2,1], [2,2], [1,2], [1,1]]) error")
	}
}

func TestCreatePolygonZ(t *testing.T) {
	shell := []CoordZ{
		CoordZ{0, 0, 0},
		CoordZ{3, 0, 0},
		CoordZ{3, 3, 0},
		CoordZ{0, 3, 0},
		CoordZ{0, 0, 0}}

	geom := CreatePolygonZ(shell)
	if geom != nil && geom.GetType() == POLYGON {
		t.Logf("Log: CreatePolygonZ([[0,0,0], [3,0,0], [3,3,0], [0,3,0], [0,0,0]]) returns %q", geom.ToWKT())
	} else {
		t.Errorf("Error: CreatePolygonZ([[0,0,0], [3,0,0], [3,3,0], [0,3,0], [0,0,0]]) error")
	}
}

func TestCreateMultiGeometry(t *testing.T) {
	var geoms []*CGeometry
	var geom *CGeometry

	geoms = append(geoms, CreatePoint(0, 0))
	geoms = append(geoms, CreatePoint(1, 1))

	geoms = append(geoms, CreateLineString([]Point{Point{0, 0}, Point{1, 1}}))
	geoms = append(geoms, CreateLineString([]Point{Point{0, 0}, Point{-1, -1}}))

	geoms = append(geoms, CreatePolygon([]Point{Point{0, 0}, Point{3, 0}, Point{3, 3}, Point{0, 3}, Point{0, 0}}))
	geoms = append(geoms, CreatePolygon([]Point{Point{0, 0}, Point{-3, 0}, Point{-3, -3}, Point{0, -3}, Point{0, 0}}))

	geom = CreateMultiGeometry(geoms, MULTIPOINT)
	if geom != nil && geom.GetType() == MULTIPOINT {
		t.Logf("Log: CreateMultiGeometry([[0,0], [1,1]], MULTIPOINT) returns %q", geom.ToWKT())
	} else {
		t.Errorf("Error: CreateMultiGeometry([[0,0], [1,1]], MULTIPOINT) error")
	}

	geom = CreateMultiGeometry(geoms, MULTILINESTRING)
	if geom != nil && geom.GetType() == MULTILINESTRING {
		t.Logf("Log: CreateMultiGeometry([[[0,0], [1,1]], [[0,0], [-1,-1]]], MULTILINESTRING) returns %q", geom.ToWKT())
	} else {
		t.Errorf("Error: CreateMultiGeometry([[[0,0], [1,1]], [[0,0], [-1,-1]]], MULTILINESTRING) error")
	}

	geom = CreateMultiGeometry(geoms, MULTIPOLYGON)
	if geom != nil && geom.GetType() == MULTIPOLYGON {
		t.Logf("Log: CreateMultiGeometry([[[[0,0], [3,0], [3,3], [0,3], [0,0]], [[0,0], [-3,0], [-3,-3], [0,-3], [0,0]]]], MULTIPOLYGON) returns %q", geom.ToWKT())
	} else {
		t.Errorf("Error: CreateMultiGeometry([[[[0,0], [3,0], [3,3], [0,3], [0,0]], [[0,0], [-3,0], [-3,-3], [0,-3], [0,0]]]], MULTIPOLYGON) error")
	}

	geom = CreateMultiGeometry(geoms, GEOMETRYCOLLECTION)
	if geom != nil && geom.GetType() == GEOMETRYCOLLECTION {
		t.Logf("Log: CreateMultiGeometry(all, GEOMETRYCOLLECTION) returns %q", geom.ToWKT())
	} else {
		t.Errorf("Error: CreateMultiGeometry(all, GEOMETRYCOLLECTION) error")
	}
}

func TestPolygonize(t *testing.T) {
	var geoms []*CGeometry

	geoms = append(geoms, CreateFromWKT("LINESTRING(0 0, 10 0, 10 10, 0 10, 0 0)"))
	geoms = append(geoms, CreateFromWKT("LINESTRING(5 -5, 5 15)"))

	geom := Polygonize(geoms)
	if geom != nil {
		t.Logf("Log: Polygonize() returns %q", geom.ToWKT())
	} else {
		t.Errorf("Error: Polygonize() error")
	}
}

func TestClone(t *testing.T) {
	p1 := CreatePoint(116.39, 39.9)
	p2 := p1.Clone()

	if strings.EqualFold(p1.ToWKT(), p2.ToWKT()) {
		t.Logf("Log: Clone() returns correct result %q", p2.ToWKT())
	} else {
		t.Errorf("Error: Clone() error")
	}
}

func TestGeometryInfo(t *testing.T) {
	pt := CreatePoint(116.39, 39.9)

	srid := 4326
	pt.SetSRID(srid)

	if pt.GetSRID() == srid {
		t.Logf("Log: SetSRID() and GetSRID() returns correct results %d", srid)
	} else {
		t.Errorf("Error: SetSRID() or GetSRID() error")
	}

	if pt.GetNumGeometries() != 1 {
		t.Errorf("Error: GetNumGeometries() error")
	}

	if pt.GetGeometryN(0) == nil {
		t.Errorf("Error: GetGeometryN() error")
	}

	shell := []Point{
		Point{0, 0},
		Point{3, 0},
		Point{3, 3},
		Point{0, 3},
		Point{0, 0}}

	hole := []Point{
		Point{1, 1},
		Point{2, 1},
		Point{2, 2},
		Point{1, 2},
		Point{1, 1}}

	pg := CreatePolygon(shell, hole)

	if pg.GetNumInteriorRings() != 1 {
		t.Errorf("Error: GetNumInteriorRings() error")
	}

	if pg.GetExteriorRing() == nil {
		t.Errorf("Error: GetExteriorRing() error")
	}

	if pg.GetInteriorRingN(0) == nil {
		t.Errorf("Error: GetInteriorRingN() error")
	}

	if pg.GetNumCoordinates() != 10 {
		t.Errorf("Error: GetNumCoordinates() error")
	}

}

func TestGetXY(t *testing.T) {
	pt := CreatePoint(1, 2)
	x, y := pt.GetCoord().X, pt.GetCoord().Y

	if x != 1 || y != 2 {
		t.Errorf("Error: GetXY() error")
	}
}

func TestGetCoords(t *testing.T) {
	coords := []Point{
		Point{0, 0},
		Point{1, 1}}

	ln := CreateLineString(coords)
	coordsGot := ln.GetCoords()

	if coordsGot == nil {
		t.Errorf("Error: GetCoords() error")
	} else {
		for i := range coordsGot {
			if coords[i].X != coordsGot[i].X || coords[i].Y != coordsGot[i].Y {
				t.Errorf("Error: GetCoords() error")
			}
		}
	}
}

func TestArea(t *testing.T) {
	pg := CreatePolygon([]Point{
		Point{0, 0},
		Point{0, 3},
		Point{3, 3},
		Point{3, 0},
		Point{0, 0}})

	if pg.Area() != 9 {
		t.Errorf("Error: Area() error")
	}
}

func TestLength(t *testing.T) {
	ln := CreateLineString([]Point{
		Point{0, 0},
		Point{0, 10}})

	if ln.Length() != 10 {
		t.Errorf("Error: Length() error")
	}
}

func TestDistance(t *testing.T) {
	ln1 := CreateLineString([]Point{
		Point{1, 0},
		Point{1, 2}})
	ln2 := CreateLineString([]Point{
		Point{0, 0},
		Point{0, 2}})

	if ln1.Distance(ln2) != 1 {
		t.Errorf("Error: Distance() error")
	}
}

func TestHausdorffDistance(t *testing.T) {
	ln1 := CreateLineString([]Point{
		Point{0, 0},
		Point{1, 0}})
	ln2 := CreateLineString([]Point{
		Point{0, 0},
		Point{0, 2}})

	if ln1.HausdorffDistance(ln2) != 2 {
		t.Errorf("Error: HausdorffDistance() error")
	}
}

func TestHausdorffDistanceDensify(t *testing.T) {
	a := CreateFromWKT("LINESTRING (0 0, 100 0, 10 100, 10 100)")
	b := CreateFromWKT("LINESTRING (0 100, 0 10, 80 10)")

	dis1 := a.HausdorffDistanceDensify(b, 0.5)
	if dis1 != 40 {
		t.Errorf("Error: HausdorffDistance() error")
	}

	dis2 := a.HausdorffDistanceDensify(b, 1)
	if math.Abs(dis2-22) > 1 {
		t.Errorf("Error: HausdorffDistance() error")
	}
}

func TestNearestPoints(t *testing.T) {
	a := CreateFromWKT("POINT (0 0)")
	b := CreateFromWKT("LINESTRING (0 100, 100 0, 0 -100)")

	coords := a.NearestPoints(b)
	if !coords[0].Equals(&Point{0, 0}) || !coords[1].Equals(&Point{50, 50}) {
		t.Errorf("Error: NearestPoints() error")
	}
}

func TestNearestPointZs(t *testing.T) {
	a := CreateFromWKT("POINT (0 0 0)")
	b := CreateFromWKT("LINESTRING (0 0 100, 100 0 100)")

	coords := a.NearestPointZs(b)
	if !coords[0].Equals(&CoordZ{0, 0, 0}) || !coords[1].Equals(&CoordZ{0, 0, 100}) {
		t.Errorf("Error: NearestPointZs() error")
	}
}

func TestDisjoint(t *testing.T) {
	pt := CreatePoint(0, 0)
	ln := CreateLineString([]Point{
		Point{1, 1},
		Point{2, 1}})

	if !pt.Disjoint(ln) {
		t.Errorf("Error: Disjoint() error")
	}
}

func TestTouches(t *testing.T) {
	pt := CreatePoint(1, 1)
	ln := CreateLineString([]Point{
		Point{1, 1},
		Point{2, 1}})

	if !pt.Touches(ln) {
		t.Errorf("Error: Touches() error")
	}
}

func TestIntersects(t *testing.T) {
	pt := CreatePoint(1.5, 1)
	ln := CreateLineString([]Point{
		Point{1, 1},
		Point{2, 1}})

	if !pt.Intersects(ln) {
		t.Errorf("Error: Intersects() error")
	}
}

func TestCrosses(t *testing.T) {
	ln1 := CreateLineString([]Point{
		Point{1, 1},
		Point{2, 1}})
	ln2 := CreateLineString([]Point{
		Point{1.5, 0},
		Point{1.5, 2}})

	if !ln1.Crosses(ln2) {
		t.Errorf("Error: Crosses() error")
	}
}

func TestWithin(t *testing.T) {
	pt := CreatePoint(0, 0)

	if !pt.Within(pt) {
		t.Errorf("Error: Within() error")
	}
}

func TestContains(t *testing.T) {
	pt := CreatePoint(0, 0)
	ln := CreateLineString([]Point{
		Point{1, 1},
		Point{2, 1}})
	m := CreateMultiGeometry([]*CGeometry{pt, ln}, GEOMETRYCOLLECTION)

	if !m.Contains(ln) {
		t.Errorf("Error: Contains() error")
	}
}

func TestOverlaps(t *testing.T) {
	pg1 := CreatePolygon([]Point{
		Point{1, 0},
		Point{1, 3},
		Point{4, 3},
		Point{4, 0},
		Point{1, 0}})

	pg2 := CreatePolygon([]Point{
		Point{0, 0},
		Point{0, 3},
		Point{3, 3},
		Point{3, 0},
		Point{0, 0}})

	if !pg1.Overlaps(pg2) {
		t.Errorf("Error: Overlaps() error")
	}
}

func TestEquals(t *testing.T) {
	pt := CreatePoint(0, 0)

	if !pt.Equals(pt) {
		t.Errorf("Error: Equals() error")
	}
}

func TestEqualsExact(t *testing.T) {
	pt := CreatePoint(0, 0)
	pt2 := CreatePoint(0.001, 0)

	if !pt.EqualsExact(pt2, 0.01) {
		t.Errorf("Error: EqualsExact() error")
	}
}

func TestCovers(t *testing.T) {
	ln := CreateLineString([]Point{
		Point{1, 1},
		Point{2, 1}})
	pt := CreatePoint(1, 1)

	if !ln.Covers(pt) {
		t.Errorf("Error: Covers() error")
	}
}

func TestCoveredBy(t *testing.T) {
	pt := CreatePoint(1, 1)
	ln := CreateLineString([]Point{
		Point{1, 1},
		Point{2, 1}})

	if !pt.CoveredBy(ln) {
		t.Errorf("Error: CoveredBy() error")
	}
}

func TestRelatePattern(t *testing.T) {
	pt := CreatePoint(1, 1)

	if !pt.RelatePattern(pt, "T*F**FFF*") {
		t.Errorf("Error: RelatePattern() error")
	}
}

func TestRelate(t *testing.T) {
	pt := CreatePoint(1, 1)
	relate := pt.Relate(pt)

	if relate != "0FFFFFFF2" {
		t.Errorf("Error: Relate() error")
	}
}

func TestNormalize(t *testing.T) {
	var geoms1 []*CGeometry
	var geoms2 []*CGeometry

	geoms1 = append(geoms1, CreatePoint(0, 0), CreatePoint(1, 1))
	geoms2 = append(geoms2, CreatePoint(1, 1), CreatePoint(0, 0))

	geom1 := CreateMultiGeometry(geoms1, MULTIPOINT)
	geom2 := CreateMultiGeometry(geoms2, MULTIPOINT)

	geom1.Normalize()
	geom2.Normalize()

	if !geom1.Equals(geom2) {
		t.Errorf("Error: Normalize() error")
	}
}

func TestIsValid(t *testing.T) {
	shell := []Point{
		Point{0, 0},
		Point{3, 0},
		Point{3, 3},
		Point{0, 3},
		Point{0, 0}}

	hole := []Point{
		Point{1, 1},
		Point{2, 1},
		Point{2, 2},
		Point{1, 2},
		Point{1, 1}}

	pg := CreatePolygon(hole, shell)

	if pg.IsValid() {
		t.Errorf("Error: IsValid() error")
	}
}

func TestIsEmpty(t *testing.T) {
	pt := CreatePoint(1, 1)

	if pt.IsEmpty() {
		t.Errorf("Error: IsEmpty() error")
	}

	empty := CreateFromWKT("GEOMETRYCOLLECTION EMPTY")

	if !empty.IsEmpty() {
		t.Errorf("Error: IsEmpty() error")
	}
}

func TestIsSimple(t *testing.T) {
	ln := CreateLineString([]Point{
		Point{2, 1},
		Point{1, 1},
		Point{2, 1}})

	if ln.IsSimple() {
		t.Errorf("Error: IsSimple() error")
	}
}

func TestIsRing(t *testing.T) {
	ring := CreateLinearRing([]Point{
		Point{0, 0},
		Point{0, 1},
		Point{1, 1},
		Point{0, 0}})

	if !ring.IsRing() {
		t.Errorf("Error: IsRing() error")
	}
}

func TestHasZ(t *testing.T) {
	pt := CreatePointZ(1, 1, 0)

	if !pt.HasZ() {
		t.Errorf("Error: HasZ() error")
	}
}

func TestIsClosed(t *testing.T) {
	ln := CreateLineString([]Point{
		Point{2, 1},
		Point{1, 1},
		Point{2, 1}})

	if !ln.IsClosed() {
		t.Errorf("Error: IsClosed() error")
	}
}

func TestEnvelope(t *testing.T) {
	var env *CGeometry

	pt := CreatePoint(0, 0)
	env = pt.Envelope()

	if env == nil {
		t.Errorf("Error: Point.Envelope() error")
	} else {
		if env.GetType() != POINT {
			t.Errorf("Error: Point.Envelope() returns error result")
		}
	}

	ln := CreateFromWKT("LINESTRING (0 0, 100 0)")
	env = ln.Envelope()

	if env == nil {
		t.Errorf("Error: LineString.Envelope() error")
	} else {
		if env.GetType() != POLYGON {
			t.Errorf("Error: LineString.Envelope() returns error result")
		}
	}
}

func TestIntersection(t *testing.T) {
	pt := CreateFromWKT("POINT (0 0)")
	ln := CreateFromWKT("LINESTRING (0 0, 100 0)")

	geom := pt.Intersection(ln)

	if geom == nil {
		t.Errorf("Error: Intersection() error")
	} else {
		t.Logf("Log: Intersection() returns %q", geom.ToWKT())
	}
}

func TestConvexHull(t *testing.T) {
	m := CreateFromWKT("MULTIPOINT ((10 40), (40 30), (20 20), (30 10))")
	geom := m.ConvexHull()

	if geom == nil {
		t.Errorf("Error: ConvexHull() error")
	} else {
		t.Logf("Log: ConvexHull() returns %q", geom.ToWKT())
	}
}

func TestDifference(t *testing.T) {
	geom1 := CreateFromWKT("LINESTRING (0 0, 100 0)")
	geom2 := CreateFromWKT("LINESTRING (0 0, 50 0)")

	geom := geom1.Difference(geom2)

	if geom == nil {
		t.Errorf("Error: Difference() error")
	} else {
		t.Logf("Log: Difference() returns %q", geom.ToWKT())
	}
}

func TestSymDifference(t *testing.T) {
	geom1 := CreateFromWKT("LINESTRING (0 0, 100 0)")
	geom2 := CreateFromWKT("LINESTRING (0 0, 50 0)")

	geom := geom1.SymDifference(geom2)

	if geom == nil {
		t.Errorf("Error: SymDifference() error")
	} else {
		t.Logf("Log: SymDifference() returns %q", geom.ToWKT())
	}
}

func TestBoundary(t *testing.T) {
	ln := CreateFromWKT("LINESTRING (0 0, 100 0)")

	geom := ln.Boundary()

	if geom == nil {
		t.Errorf("Error: Boundary() error")
	} else {
		t.Logf("Log: Boundary() returns %q", geom.ToWKT())
	}
}

func TestUnion(t *testing.T) {
	geom1 := CreateFromWKT("POINT (0 0)")
	geom2 := CreateFromWKT("LINESTRING (5 0, 10 0)")

	geom := geom1.Union(geom2)

	if geom == nil {
		t.Errorf("Error: Union() error")
	} else {
		t.Logf("Log: Union() returns %q", geom.ToWKT())
	}
}

func TestUnaryUnion(t *testing.T) {
	m := CreateFromWKT("MULTILINESTRING((0 0, 10 0), (5 -5, 5 5))")

	geom := m.UnaryUnion()

	if geom == nil {
		t.Errorf("Error: UnaryUnion() error")
	} else {
		t.Logf("Log: UnaryUnion() returns %q", geom.ToWKT())
	}
}

func TestPointOnSurface(t *testing.T) {
	pg := CreateFromWKT("POLYGON ((0 0, 10 0, 10 10, 5 10, 5 5, 0 5, 0 0))")

	geom := pg.PointOnSurface()

	if geom == nil {
		t.Errorf("Error: PointOnSurface() error")
	} else {
		t.Logf("Log: PointOnSurface() returns %q", geom.ToWKT())
	}
}

func TestCentroid(t *testing.T) {
	pg := CreateFromWKT("POLYGON ((0 0, 10 0, 10 10, 5 10, 5 5, 0 5, 0 0))")

	geom := pg.Centroid()

	if geom == nil {
		t.Errorf("Error: Centroid() error")
	} else {
		t.Logf("Log: Centroid() returns %q", geom.ToWKT())
	}
}

func TestNode(t *testing.T) {
	g := CreateFromWKT("LINESTRING (0 0, 10 0)")

	geom := g.Node()

	if geom == nil {
		t.Errorf("Error: Node() error")
	} else {
		t.Logf("Log: Node() returns %q", geom.ToWKT())
	}
}

func TestSimplify(t *testing.T) {
	g := CreateFromWKT("LINESTRING (0 0, 10 0, 10 0)")

	geom := g.Simplify(0)

	if geom == nil {
		t.Errorf("Error: Simplify() error")
	} else {
		t.Logf("Log: Simplify() returns %q", geom.ToWKT())
	}
}

func TestTopologyPreserveSimplify(t *testing.T) {
	g := CreateFromWKT("LINESTRING (0 0, 10 0.1, 20 0)")

	geom := g.TopologyPreserveSimplify(0.2)

	if geom == nil {
		t.Errorf("Error: TopologyPreserveSimplify() error")
	} else {
		t.Logf("Log: TopologyPreserveSimplify() returns %q", geom.ToWKT())
	}
}

func TestExtractUniquePoints(t *testing.T) {
	g := CreateFromWKT("LINESTRING (0 0, 10 0, 10 0)")

	geom := g.ExtractUniquePoints()

	if geom == nil {
		t.Errorf("Error: ExtractUniquePoints() error")
	} else {
		t.Logf("Log: ExtractUniquePoints() returns %q", geom.ToWKT())
	}
}

func TestSharedPaths(t *testing.T) {
	g := CreateFromWKT("LINESTRING (0 0, 10 0)")
	line := CreateFromWKT("LINESTRING (5 0, 15 0, 8 0)")

	geom := g.SharedPaths(line)

	if geom == nil {
		t.Errorf("Error: SharedPaths() error")
	} else {
		t.Logf("Log: SharedPaths() returns %q", geom.ToWKT())
	}
}

func TestSnap(t *testing.T) {
	g := CreateFromWKT("POINT (0 0)")
	line := CreateFromWKT("LINESTRING (0 -1, 10 0)")

	geom := g.Snap(line, 2)

	if geom == nil {
		t.Errorf("Error: Snap() error")
	} else {
		t.Logf("Log: Snap() returns %q", geom.ToWKT())
	}
}

func TestDelaunayTriangulation(t *testing.T) {
	g := CreateFromWKT("LINESTRING (0 0, 10 0, 10 10)")

	geom := g.DelaunayTriangulation(1, false)

	if geom == nil {
		t.Errorf("Error: DelaunayTriangulation() error")
	} else {
		t.Logf("Log: DelaunayTriangulation() returns %q", geom.ToWKT())
	}
}

func TestBuffer(t *testing.T) {
	pt := CreatePoint(0, 0)
	var geom *CGeometry

	width := 10.0
	quadsegs := 8

	geom = pt.Buffer(width)
	if geom == nil {
		t.Errorf("Error: Buffer() error")
	} else {
		ptCount := geom.GetNumCoordinates()
		if ptCount != 4*quadsegs+1 {
			t.Errorf("Error: Buffer() returns error polygon")
		} else {
			t.Logf("Log: Buffer([0, 0]) returns %q", geom.ToWKT())
		}
	}

	geom = pt.BufferWithStyle(width, quadsegs, CAP_ROUND, JOIN_ROUND, 5)
	if geom == nil {
		t.Errorf("Error: BufferWithStyle() error")
	} else {
		ptCount := geom.GetNumCoordinates()
		if ptCount != 4*quadsegs+1 {
			t.Errorf("Error: BufferWithStyle() returns error polygon")
		} else {
			t.Logf("Log: BufferWithStyle([0, 0]) returns %q", geom.ToWKT())
		}
	}
}

func TestOffsetCurve(t *testing.T) {
	ln := CreateLineString([]Point{
		Point{0, 0},
		Point{10, 0},
		Point{10, 10}})

	width := -10.0
	quadsegs := 8

	geom := ln.OffsetCurve(width, quadsegs, JOIN_MITRE, 5)

	if geom == nil {
		t.Errorf("Error: OffsetCurve() error")
	} else {
		t.Logf("Log: OffsetCurve() returns %q", geom.ToWKT())
	}
}

func TestProject(t *testing.T) {
	ln := CreateFromWKT("LINESTRING (0 0, 10 0, 10 10)")
	pt := CreateFromWKT("POINT (5 2)")

	if ln.Project(pt) != 5 {
		t.Errorf("Error: Project() error")
	}

	if ln.ProjectNormalized(pt) != 0.25 {
		t.Errorf("Error: Project() error")
	}
}

func TestInterpolate(t *testing.T) {
	ln := CreateFromWKT("LINESTRING (0 0, 10 0, 10 10)")

	pt1 := ln.Interpolate(10)
	if pt1 == nil || !pt1.Equals(CreatePoint(10, 0)) {
		t.Errorf("Error: Interpolate() error")
	}

	pt2 := ln.InterpolateNormalized(0.25)
	if pt2 == nil || !pt2.Equals(CreatePoint(5, 0)) {
		t.Errorf("Error: InterpolateNormalized() error")
	}
}
