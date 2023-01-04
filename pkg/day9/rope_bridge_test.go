package day9

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
			args{`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`},
			13,
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
		{"result for solution1", 5710},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution1(); got != tt.want {
				t.Errorf("Solution1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doSolution2(t *testing.T) {
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
			args{`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`},
			1,
		},
		{
			"it should return the provided second example result",
			args{`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`},
			36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doSolution2(tt.args.raw); got != tt.want {
				t.Errorf("doSolution2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"result for solution2", 2259},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution2(); got != tt.want {
				t.Errorf("Solution2() = %v, want %v", got, tt.want)
			}
		})
	}
}
