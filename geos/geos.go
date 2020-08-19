/*
package geo implements a wrapper for GEOS library.
*/
package geos

/*
#cgo CFLAGS: -I${SRCDIR}/lib/include
#cgo darwin LDFLAGS: -L${SRCDIR}/lib/darwin -lgeos_c
#cgo linux LDFLAGS: -L${SRCDIR}/lib/linux -lgeos_c
#include <geos_c.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>

void go_notice_handler(const char *fmt, ...) {
    va_list ap;
    va_start(ap, fmt);
    fprintf(stdout, "NOTICE: ");
    vfprintf(stdout, fmt, ap);
    va_end(ap);
}

void go_error_handler(const char *fmt, ...) {
    va_list ap;
    va_start(ap, fmt);
    fprintf(stderr, "Error: ");
    vfprintf(stderr, fmt, ap);
    va_end(ap);
}

GEOSContextHandle_t go_initGEOS_r() {
    return initGEOS_r(go_notice_handler, go_error_handler);
}

*/
import "C"

func init() {
	ctxHandle = C.go_initGEOS_r()
}

const (
	GEOS_CAPI_VERSION_MAJOR int    = C.GEOS_CAPI_VERSION_MAJOR
	GEOS_CAPI_VERSION_MINOR int    = C.GEOS_CAPI_VERSION_MINOR
	GEOS_CAPI_VERSION_PATCH int    = C.GEOS_CAPI_VERSION_PATCH
	GEOS_CAPI_VERSION       string = C.GEOS_CAPI_VERSION
	GEOS_VERSION_MAJOR      int    = C.GEOS_VERSION_MAJOR
	GEOS_VERSION_MINOR      int    = C.GEOS_VERSION_MINOR
	GEOS_VERSION_PATCH      int    = C.GEOS_VERSION_PATCH
	GEOS_VERSION            string = C.GEOS_VERSION
	GEOS_JTS_PORT           string = C.GEOS_JTS_PORT
)

var (
	ctxHandle C.GEOSContextHandle_t
)

func FinishGEOS() {
	C.finishGEOS_r(ctxHandle)
}

func Version() string {
	return C.GoString(C.GEOSversion())
}
