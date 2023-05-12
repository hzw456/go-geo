package geo

import (
	"math"
)

func GetAzimuth(p1, p2 Point, srid SRID) float64 {
	if srid != SRID_WGS84_GPS {
		return 0
	}

	lon1_rad := p1.X * math.Pi / 180
	lon2_rad := p2.X * math.Pi / 180
	lat1_rad := p1.Y * math.Pi / 180
	lat2_rad := p2.Y * math.Pi / 180

	y := math.Sin(lon2_rad-lon1_rad) * math.Cos(lat2_rad)
	x := math.Cos(lat1_rad)*math.Sin(lat2_rad) - math.Sin(lat1_rad)*math.Cos(lat2_rad)*math.Cos(lon2_rad-lon1_rad)

	brng := math.Atan2(y, x) * 180 / math.Pi

	return math.Mod(brng+360, 360)
}
