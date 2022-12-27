package day6

import "testing"

func Test_findMarker(t *testing.T) {
	type args struct {
		stream     string
		markerSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it should validate first example",
			args{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4},
			7,
		},
		{
			"it should validate second example",
			args{"bvwbjplbgvbhsrlpgdmjqwftvncz", 4},
			5,
		},
		{
			"it should validate third example",
			args{"nppdvjthqldpwncqszvftbrmjlhg", 4},
			6,
		},
		{
			"it should validate fourth example",
			args{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4},
			10,
		},
		{
			"it should validate fifth example",
			args{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMarker(tt.args.stream, tt.args.markerSize); got != tt.want {
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

func TestSolution2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"result for solution2", 3645},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution2(); got != tt.want {
				t.Errorf("Solution2() = %v, want %v", got, tt.want)
			}
		})
	}
}
