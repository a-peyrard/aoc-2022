package day3

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
			args{`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`},
			157,
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

func Test_priority(t *testing.T) {
	type args struct {
		item byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it should return correct priority for 'P'",
			args{'P'},
			42,
		},
		{
			"it should return correct priority for 't'",
			args{'t'},
			20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := priority(tt.args.item); got != tt.want {
				t.Errorf("priority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"result for solution1", 8072},
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
			args{`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`},
			70,
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
		{"result for solution1", 2567},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution2(); got != tt.want {
				t.Errorf("Solution2() = %v, want %v", got, tt.want)
			}
		})
	}
}
