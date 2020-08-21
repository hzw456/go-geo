package geos

import "runtime"

/*
#cgo darwin CFLAGS: -I./ -I${SRCDIR}/darwin/include
#cgo darwin CPPFLAGS: -I./ -I${SRCDIR}/darwin/include
#cgo darwin LDFLAGS: -L./ -L${SRCDIR}/darwin/lib -lgeos_c
#cgo linux CFLAGS: -I./ -I${SRCDIR}/linux/include
#cgo linux CPPFLAGS: -I./ -I${SRCDIR}/linux/include
#cgo linux LDFLAGS:  -L./ -L${SRCDIR}/linux/lib -lgeos_c
#cgo CXXFLAGS: --std=c++11
#include <geos_c.h>

*/
import "C"

type CGeometryType int

const (
	POINT              CGeometryType = C.GEOS_POINT
	LINESTRING         CGeometryType = C.GEOS_LINESTRING
	LINEARRING         CGeometryType = C.GEOS_LINEARRING
	POLYGON            CGeometryType = C.GEOS_POLYGON
	MULTIPOINT         CGeometryType = C.GEOS_MULTIPOINT
	MULTILINESTRING    CGeometryType = C.GEOS_MULTILINESTRING
	MULTIPOLYGON       CGeometryType = C.GEOS_MULTIPOLYGON
	GEOMETRYCOLLECTION CGeometryType = C.GEOS_GEOMETRYCOLLECTION
)

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

// Only support Point
func (g *CGeometry) GetCoord() Point {
	var x, y C.double
	C.GEOSGeomGetX_r(ctxHandle, g.c, &x)
	C.GEOSGeomGetY_r(ctxHandle, g.c, &y)
	return Point{float64(x), float64(y)}
}

// Only support LineString, LinearRing or Point
func (g *CGeometry) GetCoords() []Point {
	c := C.GEOSGeom_getCoordSeq_r(ctxHandle, g.c)
	coordSeq := coordSeqFromC(c, false)
	return coordSeq.toCoords()
}

// Only support Polygon
func (g *CGeometry) GetCoordsSlice() [][]Point {
	var coords [][]Point
	coords = append(coords, g.GetExteriorRing().GetCoords())
	for i := 0; i < g.GetNumInteriorRings(); i++ {
		coords = append(coords, g.GetInteriorRingN(i).GetCoords())
	}
	return coords
}

// Only support LineString, LinearRing or Point
func (g *CGeometry) GetCoordZs() []CoordZ {
	c := C.GEOSGeom_getCoordSeq_r(ctxHandle, g.c)
	coordSeq := coordSeqFromC(c, false)
	return coordSeq.toCoordZs()
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
