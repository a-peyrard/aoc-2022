package day1

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"it should parse a basic number",
			args{`200`},
			[][]int{{200}},
		},
		{
			"it should parse a single elf calories",
			args{`200
400`},
			[][]int{{200, 400}},
		},
		{
			"it should parse a multiple elf calories",
			args{`200
400

300
500`},
			[][]int{{200, 400}, {300, 500}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMaxCalories(t *testing.T) {
	type args struct {
		elvesPkg [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it should return the max with only one elf",
			args{[][]int{{100, 200}}},
			300,
		},
		{
			"it should return the max with multiple elves",
			args{[][]int{{100, 200}, {500}, {100, 200, 400}}},
			700,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxCalories(tt.args.elvesPkg); got != tt.want {
				t.Errorf("FindMaxCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			args{`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`},
			24000,
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
		{"result for solution1", 74711},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution1(); got != tt.want {
				t.Errorf("Solution1() = %v, want %v", got, tt.want)
			}
		})
	}
}
