package day6

import "testing"

func Test_findMarker(t *testing.T) {
	type args struct {
		stream string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it should validate first example",
			args{"mjqjpqmgbljsphdztnvjfqwrcgsmlb"},
			7,
		},
		{
			"it should validate second example",
			args{"bvwbjplbgvbhsrlpgdmjqwftvncz"},
			5,
		},
		{
			"it should validate third example",
			args{"nppdvjthqldpwncqszvftbrmjlhg"},
			6,
		},
		{
			"it should validate fourth example",
			args{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"},
			10,
		},
		{
			"it should validate fifth example",
			args{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMarker(tt.args.stream); got != tt.want {
				t.Errorf("findMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"result for solution1", 1080},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution1(); got != tt.want {
				t.Errorf("Solution1() = %v, want %v", got, tt.want)
			}
		})
	}
}
