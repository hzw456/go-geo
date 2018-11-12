package hexagon

import gogeo "github.com/sadnessly/go-geo"

type CoordUnit struct {
	Lng float64 `thrift:"lng,1,required" db:"lng" json:"lng"`
	Lat float64 `thrift:"lat,2,required" db:"lat" json:"lat"`
}

type Hexagon []*CoordUnit

// type Hexagon_s

func NewHexagon(coords ...*CoordUnit) (hex *Hexagon) {
	for _, v := range coords {
		*hex = append(*hex, v)
	}
	return hex
}

func (hex *Hexagon) toPoly() gogeo.Polygon {
	line := gogeo.LineString{}
	for _, v := range *hex {
		line.AppendPoint(gogeo.Point{v.Lng, v.Lat})
	}
	return *gogeo.NewPolygon(*gogeo.NewLinearRing(line))
}

// func (hex *Hexagon) getCenter() CoordUnit {

// }
