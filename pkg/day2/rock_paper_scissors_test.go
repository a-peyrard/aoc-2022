package day2

import "testing"

func TestDoSolution1(t *testing.T) {
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
			args{`A Y
B X
C Z`},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoSolution1(tt.args.raw); got != tt.want {
				t.Errorf("DoSolution1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"result for solution1", 10994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution1(); got != tt.want {
				t.Errorf("Solution1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoSolution2(t *testing.T) {
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
			args{`A Y
B X
C Z`},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoSolution2(tt.args.raw); got != tt.want {
				t.Errorf("DoSolution2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"result for solution1", 12526},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution2(); got != tt.want {
				t.Errorf("Solution2() = %v, want %v", got, tt.want)
			}
		})
	}
}
