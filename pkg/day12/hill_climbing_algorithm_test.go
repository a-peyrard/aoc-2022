package day12

import "testing"

const InputExample = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

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
			args{InputExample},
			31,
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
		{"result for solution1", 437},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution1(); got != tt.want {
				t.Errorf("Solution1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isWalkable(t *testing.T) {
	type args struct {
		origin byte
		dest   byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"it should walk from 'k' to 'l'",
			args{origin: 'k', dest: 'l'},
			true,
		},
		{
			"it should walk from 'k' to 'i'",
			args{origin: 'k', dest: 'i'},
			true,
		},
		{
			"it should walk from 'z' to 'a'",
			args{origin: 'z', dest: 'a'},
			true,
		},
		{
			"it should not walk from 'a' to 'c'",
			args{origin: 'a', dest: 'c'},
			false,
		},
		{
			"it should walk from 'k' to 'k'",
			args{origin: 'k', dest: 'k'},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWalkable(tt.args.origin, tt.args.dest); got != tt.want {
				t.Errorf("isWalkable() = %v, want %v", got, tt.want)
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
			args{InputExample},
			29,
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
		{"result for solution2", 430},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution2(); got != tt.want {
				t.Errorf("Solution2() = %v, want %v", got, tt.want)
			}
		})
	}
}
