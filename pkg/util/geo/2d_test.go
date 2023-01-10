package geo

import "testing"

func TestCoordinate_manhattanDistance(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		to *Coordinate
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			"it should compute manhattan distance between two coordinates",
			fields{X: 1, Y: 2},
			args{to: &Coordinate{5, 5}},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Coordinate{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := c.ManhattanDistance(*tt.args.to); got != tt.want {
				t.Errorf("manhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
