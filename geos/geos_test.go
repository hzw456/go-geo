package geos

import (
	"regexp"
	"testing"
)

func TestVersionConsts(t *testing.T) {
	t.Logf("Log: Version=%q", GEOS_VERSION)

	if GEOS_CAPI_VERSION_MAJOR <= 0 {
		t.Errorf("Error: GEOS_CAPI_VERSION_MAJOR=%d", GEOS_CAPI_VERSION_MAJOR)
	}

	if GEOS_CAPI_VERSION_MINOR <= 0 {
		t.Errorf("Error: GEOS_CAPI_VERSION_MINOR=%d", GEOS_CAPI_VERSION_MINOR)
	}

	if GEOS_CAPI_VERSION_PATCH <= 0 {
		t.Errorf("Error: GEOS_CAPI_VERSION_PATCH=%d", GEOS_CAPI_VERSION_PATCH)
	}

	if GEOS_VERSION_MAJOR <= 0 {
		t.Errorf("Error: GEOS_VERSION_MAJOR=%d", GEOS_VERSION_MAJOR)
	}

	if GEOS_VERSION_MINOR <= 0 {
		t.Errorf("Error: GEOS_VERSION_MINOR=%d", GEOS_VERSION_MINOR)
	}

	if GEOS_VERSION_PATCH <= 0 {
		t.Errorf("Error: GEOS_CAPI_VERSION_PATCH=%d", GEOS_VERSION_PATCH)
	}

	matched, err := regexp.MatchString(`^1\.12\.\d+$`, GEOS_JTS_PORT)
	if !matched || err != nil {
		t.Errorf("Error: GEOS_JTS_PORT=%q", GEOS_JTS_PORT)
	}
}

func TestVersion(t *testing.T) {
	version := Version()
	matched, err := regexp.MatchString(`^3\.\d+\.\d+-CAPI-\d+\.\d+\.\d+.+$`, version)
	if !matched || err != nil {
		t.Errorf("Error: Version() returns %q", version)
	}
}
