package tests

import (
	"testing"

	"github.com/hzw456/go-geo"
)

func TestUTM(t *testing.T) {
	t.Log(geo.FromLatLon(20, 120, true))
}
