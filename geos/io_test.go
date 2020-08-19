package geos

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestWkt(t *testing.T) {
	reader := createWktReader()
	writer := createWktWriter()

	pt := CreatePoint(116.39, 39.9)
	ptWkt := writer.write(pt)
	pt2 := reader.read(ptWkt)
	if pt2 == nil {
		t.Errorf("Error: Read and write WKT returns nil")
	} else {
		t.Logf("Log: Read and write WKT %q succeed", ptWkt)
	}

}

func TestWkb(t *testing.T) {
	reader := createWkbReader()
	writer := createWkbWriter()

	pt := CreatePoint(116.39, 39.9)
	ptWkb := writer.write(pt)
	ptWktHexWithGo := hex.EncodeToString(ptWkb)
	ptWkbHex := writer.writeHex(pt)
	ptWkbHexStr := string(ptWkbHex)
	pt2 := reader.read(ptWkb)
	pt3 := reader.readHex(ptWkbHex)
	if pt2 == nil || pt3 == nil {
		t.Errorf("Error: Read and write WKB returns nil")
	} else {
		if !strings.EqualFold(ptWktHexWithGo, ptWkbHexStr) {
			t.Errorf("Error: Read and wirte WKB with different hex result (encode in Go vs. encode in GEOS): \n %q \n %q ",
				ptWktHexWithGo, ptWkbHexStr)
		} else {
			t.Logf("Log: Read and write WKB %q succeed", ptWktHexWithGo)
			t.Logf("Log: Read and write WKBHex %q succeed", ptWkbHexStr)
		}
	}

}
