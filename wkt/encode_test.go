package wkt

import (
	"testing"

	"github.com/hzw456/go-geo"
)

func TestEncode(t *testing.T) {
	type args struct {
		geom geo.Geometry
	}
	pointzs := []geo.PointZ{
		{1, 1, 1},
		{0, 0, 1},
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 1},
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "point z",
			args: args{
				geom: *geo.NewPointZ(1, 1, 1),
			},
			want: "POINT Z (1 1 1)",
		},
		{
			name: "polygon z",
			args: args{
				geom: *geo.NewPolygonZFromPois(pointzs...),
			},
			want: "POLYGON Z ((1 1 1,0 0 1,0 0 0,1 1 0,1 1 1))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.geom); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
