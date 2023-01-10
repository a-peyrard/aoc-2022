package interval

import (
	"reflect"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	type args struct {
		intervals []*Interval
	}
	tests := []struct {
		name string
		args args
		want []*Interval
	}{
		{
			"it should merge intervals example 1",
			args{[]*Interval{
				CreateInterval(1, 3),
				CreateInterval(2, 4),
				CreateInterval(6, 8),
				CreateInterval(9, 10),
			}},
			[]*Interval{
				CreateInterval(1, 4),
				CreateInterval(6, 8),
				CreateInterval(9, 10),
			},
		},
		{
			"it should merge intervals example 2",
			args{[]*Interval{
				CreateInterval(6, 8),
				CreateInterval(1, 9),
				CreateInterval(2, 4),
				CreateInterval(4, 7),
			}},
			[]*Interval{
				CreateInterval(1, 9),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeIntervals(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_Length(t *testing.T) {
	tests := []struct {
		name string
		i    Interval
		want int
	}{
		{
			"it should calculate length for an interval",
			*CreateInterval(1, 5),
			5,
		},
		{
			"it should calculate length for a negative interval",
			*CreateInterval(-5, -2),
			4,
		},
		{
			"it should calculate length for a minimum interval",
			*CreateInterval(3, 3),
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}
