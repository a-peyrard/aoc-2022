package day17

import "testing"

func Test_doSolution1(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it should return the provided example result",
			args{">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"},
			3068,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doSolution1(tt.args.raw); got != tt.want {
				t.Errorf("doSolution1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"result for solution1", 3179},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution1(); got != tt.want {
				t.Errorf("Solution1() = %v, want %v", got, tt.want)
			}
		})
	}
}
