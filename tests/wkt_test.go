package tests

import (
	"testing"

	"github.com/hzw456/go-geo/wkt"
)

func TestWktDecode(t *testing.T) {
	wkt.Decode("MULTIPOLYGON (((30 20, 45 40, 10 40, 30 20)),((15 5, 40 10, 10 20, 5 10, 15 5)))")
}
