package geos

import (
	"runtime"
)

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

type coordSequence struct {
	c *C.GEOSCoordSequence
}

func (s *coordSequence) getSize() int {
	var size C.uint
	C.GEOSCoordSeq_getSize_r(ctxHandle, s.c, &size)

	return int(size)
}

func (s *coordSequence) setX(idx int, val float64) {
	C.GEOSCoordSeq_setX_r(ctxHandle, s.c, C.uint(idx), C.double(val))
}

func (s *coordSequence) setY(idx int, val float64) {
	C.GEOSCoordSeq_setY_r(ctxHandle, s.c, C.uint(idx), C.double(val))
}

func (s *coordSequence) setZ(idx int, val float64) {
	C.GEOSCoordSeq_setZ_r(ctxHandle, s.c, C.uint(idx), C.double(val))
}

func (s *coordSequence) getX(idx int) float64 {
	var val C.double
	i := C.GEOSCoordSeq_getX_r(ctxHandle, s.c, C.uint(idx), &val)
	if i == 0 {
		return 0.0
	}

	return float64(val)
}

func (s *coordSequence) getY(idx int) float64 {
	var val C.double
	i := C.GEOSCoordSeq_getY_r(ctxHandle, s.c, C.uint(idx), &val)
	if i == 0 {
		return 0.0
	}

	return float64(val)
}

func (s *coordSequence) getZ(idx int) float64 {
	var val C.double
	i := C.GEOSCoordSeq_getZ_r(ctxHandle, s.c, C.uint(idx), &val)
	if i == 0 {
		return 0.0
	}

	return float64(val)
}

func (s *coordSequence) toCoords() []Point {
	var coords []Point

	count := s.getSize()
	for i := 0; i < count; i++ {
		coord := Point{s.getX(i), s.getY(i)}
		coords = append(coords, coord)
	}

	return coords
}

func (s *coordSequence) toCoordZs() []CoordZ {
	var coords []CoordZ

	count := s.getSize()
	for i := 0; i < count; i++ {
		coord := CoordZ{s.getX(i), s.getY(i), s.getZ(i)}
		coords = append(coords, coord)
	}

	return coords
}

func coordSeqFromC(c *C.GEOSCoordSequence, hasOwnership bool) *coordSequence {
	if c == nil {
		return nil
	}

	coordSeq := &coordSequence{c: c}

	if hasOwnership {
		runtime.SetFinalizer(coordSeq, func(cs *coordSequence) {
			C.GEOSCoordSeq_destroy_r(ctxHandle, cs.c)
		})
	}

	return coordSeq
}

func coordSeqFromCoords(coords []Point, hasOwnership bool) *coordSequence {
	size := len(coords)
	coordSeq := createCoordSeq(size, 2, hasOwnership)

	for i := 0; i < size; i++ {
		coord := coords[i]
		coordSeq.setX(i, coord.X)
		coordSeq.setY(i, coord.Y)
	}

	return coordSeq
}

func coordSeqFromCoordZs(coords []CoordZ, hasOwnership bool) *coordSequence {
	size := len(coords)
	coordSeq := createCoordSeq(size, 3, hasOwnership)

	for i := 0; i < size; i++ {
		coord := coords[i]
		coordSeq.setX(i, coord.X)
		coordSeq.setY(i, coord.Y)
		coordSeq.setZ(i, coord.Z)
	}

	return coordSeq
}

func createCoordSeq(size, dims int, hasOwnership bool) *coordSequence {
	c := C.GEOSCoordSeq_create_r(ctxHandle, C.uint(size), C.uint(dims))
	if c == nil {
		return nil
	}

	return coordSeqFromC(c, hasOwnership)
}
