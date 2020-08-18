package geo

import (
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

type wktReader struct {
	c *C.GEOSWKTReader
}

type wktWriter struct {
	c *C.GEOSWKTWriter
}

type wkbReader struct {
	c *C.GEOSWKBReader
}

type wkbWriter struct {
	c *C.GEOSWKBWriter
}

func (r *wktReader) read(wkt string) *CGeometry {
	cs := C.CString(wkt)
	defer C.free(unsafe.Pointer(cs))

	c := C.GEOSWKTReader_read_r(ctxHandle, r.c, cs)
	return geomFromC(c, true)
}

func (w *wktWriter) write(g *CGeometry) string {
	dims := C.GEOSGeom_getCoordinateDimension_r(ctxHandle, g.c)
	C.GEOSWKTWriter_setOutputDimension_r(ctxHandle, w.c, dims)
	return C.GoString(C.GEOSWKTWriter_write_r(ctxHandle, w.c, g.c))
}

func (r *wkbReader) read(wkb []byte) *CGeometry {
	var cwkb []C.uchar
	for i := range wkb {
		cwkb = append(cwkb, C.uchar(wkb[i]))
	}

	c := C.GEOSWKBReader_read_r(ctxHandle, r.c, &cwkb[0], C.size_t(len(wkb)))
	return geomFromC(c, true)
}

func (r *wkbReader) readHex(wkb []byte) *CGeometry {
	var cwkb []C.uchar
	for i := range wkb {
		cwkb = append(cwkb, C.uchar(wkb[i]))
	}

	c := C.GEOSWKBReader_readHEX_r(ctxHandle, r.c, &cwkb[0], C.size_t(len(wkb)))
	return geomFromC(c, true)
}

func (w *wkbWriter) write(g *CGeometry) []byte {
	dims := C.GEOSGeom_getCoordinateDimension_r(ctxHandle, g.c)
	C.GEOSWKBWriter_setOutputDimension_r(ctxHandle, w.c, dims)

	var size C.size_t
	var cwkb *C.uchar = C.GEOSWKBWriter_write_r(ctxHandle, w.c, g.c, &size)
	if cwkb == nil {
		return nil
	}

	len := int(size)

	ptr := unsafe.Pointer(cwkb)
	defer C.free(ptr)

	var wkb []byte
	for i := 0; i < len; i++ {
		el := unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(C.uchar(0))*uintptr(i))
		wkb = append(wkb, byte(*(*C.uchar)(el)))
	}

	return wkb
}

func (w *wkbWriter) writeHex(g *CGeometry) []byte {
	dims := C.GEOSGeom_getCoordinateDimension_r(ctxHandle, g.c)
	C.GEOSWKBWriter_setOutputDimension_r(ctxHandle, w.c, dims)

	var size C.size_t
	var cwkb *C.uchar = C.GEOSWKBWriter_writeHEX_r(ctxHandle, w.c, g.c, &size)
	if cwkb == nil {
		return nil
	}

	len := int(size)

	ptr := unsafe.Pointer(cwkb)
	defer C.free(ptr)

	var wkb []byte
	for i := 0; i < len; i++ {
		el := unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(C.uchar(0))*uintptr(i))
		wkb = append(wkb, byte(*(*C.uchar)(el)))
	}

	return wkb
}

func createWktReader() *wktReader {
	c := C.GEOSWKTReader_create_r(ctxHandle)
	if c == nil {
		return nil
	}

	r := &wktReader{c: c}
	runtime.SetFinalizer(r, func(r *wktReader) {
		C.GEOSWKTReader_destroy_r(ctxHandle, r.c)
	})

	return r
}

func createWktWriter() *wktWriter {
	c := C.GEOSWKTWriter_create_r(ctxHandle)
	if c == nil {
		return nil
	}

	w := &wktWriter{c: c}
	runtime.SetFinalizer(w, func(w *wktWriter) {
		C.GEOSWKTWriter_destroy_r(ctxHandle, w.c)
	})

	return w
}

func createWkbReader() *wkbReader {
	c := C.GEOSWKBReader_create_r(ctxHandle)
	if c == nil {
		return nil
	}

	r := &wkbReader{c: c}
	runtime.SetFinalizer(r, func(r *wkbReader) {
		C.GEOSWKBReader_destroy_r(ctxHandle, r.c)
	})

	return r
}

func createWkbWriter() *wkbWriter {
	c := C.GEOSWKBWriter_create_r(ctxHandle)
	if c == nil {
		return nil
	}

	w := &wkbWriter{c: c}
	runtime.SetFinalizer(w, func(w *wkbWriter) {
		C.GEOSWKBWriter_destroy_r(ctxHandle, w.c)
	})

	return w
}
