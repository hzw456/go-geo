package geo

import (
	"github.com/mmcloughlin/geohash"
)

func GeohashEncode(pt Point, Precision uint, srid SRID) string {
	return geohash.EncodeWithPrecision(pt.Y, pt.X, Precision)
}

func GeohashGetCenter(hash string, srid SRID) Point {
	lat, lng := geohash.DecodeCenter(hash)
	return Point{lng, lat}
}

func GeohashGetBoundary(hash string, srid SRID) Box {
	box := geohash.BoundingBox(hash)
	return Box{MinX: box.MinLng, MinY: box.MinLat, MaxX: box.MaxLng, MaxY: box.MaxLat}
}

// N,NE,E,SE,S,SW,W,NW
func GeohashGetAllNeighbors(hash string, srid SRID) []string {
	neighbors := geohash.Neighbors(hash)
	return neighbors
}
