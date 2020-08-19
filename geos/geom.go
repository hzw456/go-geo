package geos

import (
	"math"
	"unsafe"
)

/*
#cgo CFLAGS: -I${SRCDIR}/lib/include
#cgo darwin LDFLAGS: -L${SRCDIR}/lib/darwin -lgeos_c
#cgo linux LDFLAGS: -L${SRCDIR}/lib/linux -lgeos_c
#include <geos_c.h>
#include <stdlib.h>
*/
import "C"

const (
	CAP_ROUND  CapType = C.GEOSBUF_CAP_ROUND
	CAP_FLAT   CapType = C.GEOSBUF_CAP_FLAT
	CAP_SQUARE CapType = C.GEOSBUF_CAP_SQUARE

	JOIN_ROUND JoinType = C.GEOSBUF_JOIN_ROUND
	JOIN_MITRE JoinType = C.GEOSBUF_JOIN_MITRE
	JOIN_BEVEL JoinType = C.GEOSBUF_JOIN_BEVEL
)

type CapType int
type JoinType int

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

func (g *CGeometry) GetNumCoordinates() int {
	return int(C.GEOSGetNumCoordinates_r(ctxHandle, g.c))
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
