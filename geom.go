package geo

import (
	"math"
	"runtime"
	"unsafe"
)

/*
#cgo CFLAGS: -I${SRCDIR}/geos/darwin/include
#cgo darwin LDFLAGS: -L${SRCDIR}/geos/darwin/lib -lgeos_c
#cgo linux LDFLAGS: -L${SRCDIR}/geos/linux/lib -lgeos_c
#include <geos_c.h>
#include <stdlib.h>
*/
import "C"

const (
	POINT              CGeometryType = C.GEOS_POINT
	LINESTRING         CGeometryType = C.GEOS_LINESTRING
	LINEARRING         CGeometryType = C.GEOS_LINEARRING
	POLYGON            CGeometryType = C.GEOS_POLYGON
	MULTIPOINT         CGeometryType = C.GEOS_MULTIPOINT
	MULTILINESTRING    CGeometryType = C.GEOS_MULTILINESTRING
	MULTIPOLYGON       CGeometryType = C.GEOS_MULTIPOLYGON
	GEOMETRYCOLLECTION CGeometryType = C.GEOS_GEOMETRYCOLLECTION

	CAP_ROUND  CapType = C.GEOSBUF_CAP_ROUND
	CAP_FLAT   CapType = C.GEOSBUF_CAP_FLAT
	CAP_SQUARE CapType = C.GEOSBUF_CAP_SQUARE

	JOIN_ROUND JoinType = C.GEOSBUF_JOIN_ROUND
	JOIN_MITRE JoinType = C.GEOSBUF_JOIN_MITRE
	JOIN_BEVEL JoinType = C.GEOSBUF_JOIN_BEVEL
)

type CGeometryType int

type CapType int
type JoinType int

func (t CGeometryType) ToString() string {
	switch t {
	case POINT:
		return "Point"
	case LINESTRING:
		return "LineString"
	case LINEARRING:
		return "LinearRing"
	case POLYGON:
		return "Polygon"
	case MULTIPOINT:
		return "MultiPoint"
	case MULTILINESTRING:
		return "MultiLineString"
	case MULTIPOLYGON:
		return "MultiPolygon"
	case GEOMETRYCOLLECTION:
		return "GeometryCollection"
	default:
		return "Unknown"
	}
}

type CGeometry struct {
	c *C.GEOSGeometry
}

func (g *CGeometry) Clone() *CGeometry {
	c := C.GEOSGeom_clone_r(ctxHandle, g.c)
	return geomFromC(c, true)
}

func (g *CGeometry) GetType() CGeometryType {
	return CGeometryType(C.GEOSGeomTypeId_r(ctxHandle, g.c))
}

func (g *CGeometry) SetSRID(srid int) {
	C.GEOSSetSRID_r(ctxHandle, g.c, C.int(srid))
}

func (g *CGeometry) GetSRID() int {
	return int(C.GEOSGetSRID_r(ctxHandle, g.c))
}

func (g *CGeometry) GetNumGeometries() int {
	return int(C.GEOSGetNumGeometries_r(ctxHandle, g.c))
}

func (g *CGeometry) GetGeometryN(n int) *CGeometry {
	c := C.GEOSGetGeometryN_r(ctxHandle, g.c, C.int(n))
	return geomFromC(c, false)
}

// Only support Polygon
func (g *CGeometry) GetNumInteriorRings() int {
	return int(C.GEOSGetNumInteriorRings_r(ctxHandle, g.c))
}

// Only support Polygon
func (g *CGeometry) GetInteriorRingN(n int) *CGeometry {
	c := C.GEOSGetInteriorRingN_r(ctxHandle, g.c, C.int(n))
	return geomFromC(c, false)
}

// Only support Polygon
func (g *CGeometry) GetExteriorRing() *CGeometry {
	c := C.GEOSGetExteriorRing_r(ctxHandle, g.c)
	return geomFromC(c, false)
}

func (g *CGeometry) GetNumCoordinates() int {
	return int(C.GEOSGetNumCoordinates_r(ctxHandle, g.c))
}

// Only support Point
func (g *CGeometry) GetXY() (float64, float64) {
	var x, y C.double
	C.GEOSGeomGetX_r(ctxHandle, g.c, &x)
	C.GEOSGeomGetY_r(ctxHandle, g.c, &y)

	return float64(x), float64(y)
}

// Only support LineString, LinearRing or Point
func (g *CGeometry) GetCoords() []Point {
	c := C.GEOSGeom_getCoordSeq_r(ctxHandle, g.c)
	coordSeq := coordSeqFromC(c, false)
	return coordSeq.toCoords()
}

// Only support LineString, LinearRing or Point
func (g *CGeometry) GetCoordZs() []CoordZ {
	c := C.GEOSGeom_getCoordSeq_r(ctxHandle, g.c)
	coordSeq := coordSeqFromC(c, false)
	return coordSeq.toCoordZs()
}

func (g *CGeometry) Area() float64 {
	var val C.double
	C.GEOSArea_r(ctxHandle, g.c, &val)
	return float64(val)
}

func (g *CGeometry) Length() float64 {
	var val C.double
	C.GEOSLength_r(ctxHandle, g.c, &val)
	return float64(val)
}

func (g *CGeometry) Distance(g2 *CGeometry) float64 {
	var val C.double
	C.GEOSDistance_r(ctxHandle, g.c, g2.c, &val)
	return float64(val)
}

// GEOS 3.2.0+ required
func (g *CGeometry) HausdorffDistance(g2 *CGeometry) float64 {
	var val C.double
	C.GEOSHausdorffDistance_r(ctxHandle, g.c, g2.c, &val)
	return float64(val)
}

// GEOS 3.2.0+ required
func (g *CGeometry) HausdorffDistanceDensify(g2 *CGeometry, densifyFrac float64) float64 {
	var val C.double

	if densifyFrac > 1 {
		densifyFrac = 1
	} else if densifyFrac <= 0 {
		densifyFrac = math.SmallestNonzeroFloat64
	}

	C.GEOSHausdorffDistanceDensify_r(ctxHandle, g.c, g2.c, C.double(densifyFrac), &val)
	return float64(val)
}

// GEOS 3.4.0+ required
func (g *CGeometry) NearestPoints(g2 *CGeometry) []Point {
	c := C.GEOSNearestPoints_r(ctxHandle, g.c, g2.c)
	coordSeq := coordSeqFromC(c, true)
	return coordSeq.toCoords()
}

// GEOS 3.4.0+ required
func (g *CGeometry) NearestPointZs(g2 *CGeometry) []CoordZ {
	c := C.GEOSNearestPoints_r(ctxHandle, g.c, g2.c)
	coordSeq := coordSeqFromC(c, true)
	return coordSeq.toCoordZs()
}

func (g *CGeometry) Disjoint(g2 *CGeometry) bool {
	flag := C.GEOSDisjoint_r(ctxHandle, g.c, g2.c)
	return flag == C.char(1)
}

func (g *CGeometry) Touches(g2 *CGeometry) bool {
	flag := C.GEOSTouches_r(ctxHandle, g.c, g2.c)
	return flag == C.char(1)
}

func (g *CGeometry) Intersects(g2 *CGeometry) bool {
	flag := C.GEOSIntersects_r(ctxHandle, g.c, g2.c)
	return flag == C.char(1)
}

func (g *CGeometry) Crosses(g2 *CGeometry) bool {
	flag := C.GEOSCrosses_r(ctxHandle, g.c, g2.c)
	return flag == C.char(1)
}

func (g *CGeometry) Within(g2 *CGeometry) bool {
	flag := C.GEOSWithin_r(ctxHandle, g.c, g2.c)
	return flag == C.char(1)
}

func (g *CGeometry) Contains(g2 *CGeometry) bool {
	flag := C.GEOSContains_r(ctxHandle, g.c, g2.c)
	return flag == C.char(1)
}

func (g *CGeometry) Overlaps(g2 *CGeometry) bool {
	flag := C.GEOSOverlaps_r(ctxHandle, g.c, g2.c)
	return flag == C.char(1)
}

func (g *CGeometry) Equals(g2 *CGeometry) bool {
	flag := C.GEOSEquals_r(ctxHandle, g.c, g2.c)
	return flag == C.char(1)
}

func (g *CGeometry) EqualsExact(g2 *CGeometry, tol float64) bool {
	flag := C.GEOSEqualsExact_r(ctxHandle, g.c, g2.c, C.double(tol))
	return flag == C.char(1)
}

func (g *CGeometry) Covers(g2 *CGeometry) bool {
	flag := C.GEOSCovers_r(ctxHandle, g.c, g2.c)
	return flag == C.char(1)
}

func (g *CGeometry) CoveredBy(g2 *CGeometry) bool {
	flag := C.GEOSCoveredBy_r(ctxHandle, g.c, g2.c)
	return flag == C.char(1)
}

func (g *CGeometry) RelatePattern(g2 *CGeometry, pattern string) bool {
	cs := C.CString(pattern)
	defer C.free(unsafe.Pointer(cs))

	flag := C.GEOSRelatePattern_r(ctxHandle, g.c, g2.c, cs)
	return flag == C.char(1)
}

func (g *CGeometry) Relate(g2 *CGeometry) string {
	return C.GoString(C.GEOSRelate_r(ctxHandle, g.c, g2.c))
}

func (g *CGeometry) Normalize() {
	C.GEOSNormalize_r(ctxHandle, g.c)
}

func (g *CGeometry) IsValid() bool {
	flag := C.GEOSisValid_r(ctxHandle, g.c)
	return flag == C.char(1)
}

func (g *CGeometry) IsEmpty() bool {
	flag := C.GEOSisEmpty_r(ctxHandle, g.c)
	return flag == C.char(1)
}

func (g *CGeometry) IsSimple() bool {
	flag := C.GEOSisSimple_r(ctxHandle, g.c)
	return flag == C.char(1)
}

func (g *CGeometry) IsRing() bool {
	flag := C.GEOSisRing_r(ctxHandle, g.c)
	return flag == C.char(1)
}

func (g *CGeometry) HasZ() bool {
	flag := C.GEOSHasZ_r(ctxHandle, g.c)
	return flag == C.char(1)
}

func (g *CGeometry) IsClosed() bool {
	flag := C.GEOSisClosed_r(ctxHandle, g.c)
	return flag == C.char(1)
}

func (g *CGeometry) Envelope() *CGeometry {
	c := C.GEOSEnvelope_r(ctxHandle, g.c)
	return geomFromC(c, true)
}

func (g *CGeometry) Intersection(g2 *CGeometry) *CGeometry {
	c := C.GEOSIntersection_r(ctxHandle, g.c, g2.c)
	return geomFromC(c, true)
}

func (g *CGeometry) ConvexHull() *CGeometry {
	c := C.GEOSConvexHull_r(ctxHandle, g.c)
	return geomFromC(c, true)
}

func (g *CGeometry) Difference(g2 *CGeometry) *CGeometry {
	c := C.GEOSDifference_r(ctxHandle, g.c, g2.c)
	return geomFromC(c, true)
}

func (g *CGeometry) SymDifference(g2 *CGeometry) *CGeometry {
	c := C.GEOSSymDifference_r(ctxHandle, g.c, g2.c)
	return geomFromC(c, true)
}

func (g *CGeometry) Boundary() *CGeometry {
	c := C.GEOSBoundary_r(ctxHandle, g.c)
	return geomFromC(c, true)
}

func (g *CGeometry) Union(g2 *CGeometry) *CGeometry {
	c := C.GEOSUnion_r(ctxHandle, g.c, g2.c)
	return geomFromC(c, true)
}

func (g *CGeometry) UnaryUnion() *CGeometry {
	c := C.GEOSUnaryUnion_r(ctxHandle, g.c)
	return geomFromC(c, true)
}

func (g *CGeometry) PointOnSurface() *CGeometry {
	c := C.GEOSPointOnSurface_r(ctxHandle, g.c)
	return geomFromC(c, true)
}

func (g *CGeometry) Centroid() *CGeometry {
	c := C.GEOSGetCentroid_r(ctxHandle, g.c)
	return geomFromC(c, true)
}

// GEOS 3.4.0+ required
func (g *CGeometry) Node() *CGeometry {
	c := C.GEOSNode_r(ctxHandle, g.c)
	return geomFromC(c, true)
}

func (g *CGeometry) Simplify(tol float64) *CGeometry {
	c := C.GEOSSimplify_r(ctxHandle, g.c, C.double(tol))
	return geomFromC(c, true)
}

func (g *CGeometry) TopologyPreserveSimplify(tol float64) *CGeometry {
	c := C.GEOSTopologyPreserveSimplify_r(ctxHandle, g.c, C.double(tol))
	return geomFromC(c, true)
}

func (g *CGeometry) ExtractUniquePoints() *CGeometry {
	c := C.GEOSGeom_extractUniquePoints_r(ctxHandle, g.c)
	return geomFromC(c, true)
}

// Support LineString, LinearRing only
func (g *CGeometry) SharedPaths(line *CGeometry) *CGeometry {
	c := C.GEOSSharedPaths_r(ctxHandle, g.c, line.c)
	return geomFromC(c, true)
}

// GEOS 3.3.0+ required
func (g *CGeometry) Snap(g2 *CGeometry, tol float64) *CGeometry {
	c := C.GEOSSnap_r(ctxHandle, g.c, g2.c, C.double(tol))
	return geomFromC(c, true)
}

// GEOS 3.4.0+ required
func (g *CGeometry) DelaunayTriangulation(tol float64, onlyEdges bool) *CGeometry {
	onlyEdgesC := C.int(0)
	if onlyEdges {
		onlyEdgesC = C.int(1)
	}

	c := C.GEOSDelaunayTriangulation_r(ctxHandle, g.c, C.double(tol), onlyEdgesC)
	return geomFromC(c, true)
}

func (g *CGeometry) Buffer(width float64) *CGeometry {
	c := C.GEOSBuffer_r(ctxHandle, g.c, C.double(width), C.int(8))
	return geomFromC(c, true)
}

func (g *CGeometry) BufferWithStyle(width float64, quadsegs int, endCapStyle CapType, joinStyle JoinType, mitreLimit float64) *CGeometry {
	c := C.GEOSBufferWithStyle_r(ctxHandle, g.c, C.double(width), C.int(quadsegs),
		C.int(endCapStyle), C.int(joinStyle), C.double(mitreLimit))
	return geomFromC(c, true)
}

// Only support LineString.
// Negative width for right side offset, positive width for left side offset.
func (g *CGeometry) OffsetCurve(width float64, quadsegs int, joinStyle JoinType, mitreLimit float64) *CGeometry {
	c := C.GEOSOffsetCurve_r(ctxHandle, g.c, C.double(width), C.int(quadsegs),
		C.int(joinStyle), C.double(mitreLimit))
	return geomFromC(c, true)
}

// Only support LineString.
func (g *CGeometry) Project(p *CGeometry) float64 {
	dis := C.GEOSProject_r(ctxHandle, g.c, p.c)
	return float64(dis)
}

// Only support LineString.
func (g *CGeometry) ProjectNormalized(p *CGeometry) float64 {
	dis := C.GEOSProjectNormalized_r(ctxHandle, g.c, p.c)
	return float64(dis)
}

// Only support LineString.
func (g *CGeometry) Interpolate(dis float64) *CGeometry {
	c := C.GEOSInterpolate_r(ctxHandle, g.c, C.double(dis))
	return geomFromC(c, true)
}

// Only support LineString.
func (g *CGeometry) InterpolateNormalized(dis float64) *CGeometry {
	c := C.GEOSInterpolateNormalized_r(ctxHandle, g.c, C.double(dis))
	return geomFromC(c, true)
}

func (g *CGeometry) ToWKT() string {
	writer := createWktWriter()
	return writer.write(g)
}

func (g *CGeometry) ToWKB() []byte {
	writer := createWkbWriter()
	return writer.write(g)
}

func CreateFromWKT(wkt string) *CGeometry {
	reader := createWktReader()
	return reader.read(wkt)
}

func CreateFromWKB(wkb []byte) *CGeometry {
	reader := createWkbReader()
	return reader.read(wkb)
}

func CreatePoint(x, y float64) *CGeometry {
	s := createCoordSeq(1, 2, false)
	if s == nil || s.c == nil {
		return nil
	}

	s.setX(0, x)
	s.setY(0, y)

	c := C.GEOSGeom_createPoint_r(ctxHandle, s.c)
	if c == nil {
		return nil
	}

	return geomFromC(c, true)
}

func CreatePointZ(x, y, z float64) *CGeometry {
	s := createCoordSeq(1, 3, false)
	if s == nil || s.c == nil {
		return nil
	}

	s.setX(0, x)
	s.setY(0, y)
	s.setZ(0, z)

	c := C.GEOSGeom_createPoint_r(ctxHandle, s.c)
	if c == nil {
		return nil
	}

	return geomFromC(c, true)
}

func CreateLinearRing(coords []Point) *CGeometry {
	coordSeq := coordSeqFromCoords(coords, false)
	c := C.GEOSGeom_createLinearRing_r(ctxHandle, coordSeq.c)

	return geomFromC(c, true)
}

func CreateLinearRingZ(coords []CoordZ) *CGeometry {
	coordSeq := coordSeqFromCoordZs(coords, false)
	c := C.GEOSGeom_createLinearRing_r(ctxHandle, coordSeq.c)

	return geomFromC(c, true)
}

func CreateLineString(coords []Point) *CGeometry {
	coordSeq := coordSeqFromCoords(coords, false)
	c := C.GEOSGeom_createLineString_r(ctxHandle, coordSeq.c)

	return geomFromC(c, true)
}

func CreateLineStringZ(coords []CoordZ) *CGeometry {
	coordSeq := coordSeqFromCoordZs(coords, false)
	c := C.GEOSGeom_createLineString_r(ctxHandle, coordSeq.c)

	return geomFromC(c, true)
}

func CreatePolygon(shell []Point, holes ...[]Point) *CGeometry {
	shellGeom := CreateLinearRing(shell)
	if shellGeom == nil {
		return nil
	}
	shellGeom.giveupOwnership()
	shellC := shellGeom.c

	var holesPtr **C.GEOSGeometry
	var holeCs []*C.GEOSGeometry
	for i := range holes {
		holeGeom := CreateLinearRing(holes[i])

		if holeGeom != nil {
			holeGeom.giveupOwnership()
			holeCs = append(holeCs, holeGeom.c)
		}
	}

	holeCount := len(holeCs)
	if holeCount > 0 {
		holesPtr = &holeCs[0]
	}

	c := C.GEOSGeom_createPolygon_r(ctxHandle, shellC, holesPtr, C.uint(holeCount))
	return geomFromC(c, true)
}

func CreatePolygonZ(shell []CoordZ, holes ...[]CoordZ) *CGeometry {
	shellGeom := CreateLinearRingZ(shell)
	if shellGeom == nil {
		return nil
	}
	shellGeom.giveupOwnership()
	shellC := shellGeom.c

	var holesPtr **C.GEOSGeometry
	var holeCs []*C.GEOSGeometry
	for i := range holes {
		holeGeom := CreateLinearRingZ(holes[i])

		if holeGeom != nil {
			holeGeom.giveupOwnership()
			holeCs = append(holeCs, holeGeom.c)
		}
	}

	holeCount := len(holeCs)
	if holeCount > 0 {
		holesPtr = &holeCs[0]
	}

	c := C.GEOSGeom_createPolygon_r(ctxHandle, shellC, holesPtr, C.uint(holeCount))
	return geomFromC(c, true)
}

func CreateMultiGeometry(geoms []*CGeometry, geomType CGeometryType) *CGeometry {
	var geomsCs []*C.GEOSGeometry
	for i := range geoms {
		geom := geoms[i]
		thisType := geom.GetType()
		switch {
		case geomType == MULTIPOINT && thisType != POINT:
			{
				continue
			}
		case geomType == MULTILINESTRING && thisType != LINESTRING:
			{
				continue
			}
		case geomType == MULTIPOLYGON && thisType != POLYGON:
			{
				continue
			}
		}

		geom.giveupOwnership()
		geomsCs = append(geomsCs, geom.c)
	}

	geomCount := len(geomsCs)
	if geomCount > 0 {
		geomsPtr := &geomsCs[0]
		c := C.GEOSGeom_createCollection_r(ctxHandle, C.int(geomType), geomsPtr, C.uint(geomCount))
		return geomFromC(c, true)
	}

	return nil
}

func Polygonize(geoms []*CGeometry) *CGeometry {
	var geomsCs []*C.GEOSGeometry

	for i := range geoms {
		geom := geoms[i]

		geom.giveupOwnership()
		geomsCs = append(geomsCs, geom.c)
	}

	geomCount := len(geomsCs)
	if geomCount > 0 {
		geomsPtr := &geomsCs[0]
		c := C.GEOSPolygonize_r(ctxHandle, geomsPtr, C.uint(geomCount))
		return geomFromC(c, true)
	}

	return nil
}

func geomFromC(c *C.GEOSGeometry, hasOwnership bool) *CGeometry {
	if c == nil {
		return nil
	}

	geom := &CGeometry{c: c}

	if hasOwnership {
		runtime.SetFinalizer(geom, func(g *CGeometry) {
			C.GEOSGeom_destroy_r(ctxHandle, g.c)
		})
	}

	return geom
}

// Geometry used to construct another geometry must give up its ownership.
func (g *CGeometry) giveupOwnership() {
	if g == nil {
		return
	}

	runtime.SetFinalizer(g, nil)
}
