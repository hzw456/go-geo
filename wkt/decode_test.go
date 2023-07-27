package wkt

import (
	"reflect"
	"testing"

	"github.com/hzw456/go-geo"
)

func TestDecode(t *testing.T) {
	type args struct {
		wkt string
	}
	pointzs := []geo.PointZ{
		{1, 1, 1},
		{0, 0, 1},
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 1},
	}
	tests := []struct {
		name    string
		args    args
		want    geo.Geometry
		wantErr bool
	}{
		{
			name: "point z",
			args: args{
				"POINT Z (1 1 1)",
			},
			want:    *geo.NewPointZ(1, 1, 1),
			wantErr: false,
		},
		{
			name: "point",
			args: args{
				"POINT (1 1)",
			},
			want:    *geo.NewPoint(1, 1),
			wantErr: false,
		},
		{
			name: "polygon z",
			args: args{
				"POLYGON Z ((1 1 1,0 0 1,0 0 0,1 1 0,1 1 1))",
			},
			want:    *geo.NewPolygonZFromPois(pointzs...),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.wkt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
