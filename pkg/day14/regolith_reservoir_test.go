package day14

import "testing"

const InputExample = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func Test_doSolution1(t *testing.T) {
	// fixme
	t.Skip("Skipping wip method")

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
			args{InputExample},
			24,
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
	// fixme
	t.Skip("Skipping wip method")

	tests := []struct {
		name string
		want int
	}{
		{"result for solution1", 5938},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution1(); got != tt.want {
				t.Errorf("Solution1() = %v, want %v", got, tt.want)
			}
		})
	}
}
