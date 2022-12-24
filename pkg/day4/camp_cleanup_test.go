package day4

import "testing"

func TestPair_sectionsOverlap(t *testing.T) {
	type fields struct {
		first  Section
		second Section
	}
	tests := []struct {
		name   string
		fields fields
		fully  bool
		want   bool
	}{
		{
			"it should check pair having sections with full overlaps",
			fields{
				first:  Section{start: 2, end: 8},
				second: Section{start: 3, end: 7},
			},
			true,
			true,
		},
		{
			"it should check pair having sections with full overlaps (second overlapping first)",
			fields{
				first:  Section{start: 2, end: 8},
				second: Section{start: 1, end: 9},
			},
			true,
			true,
		},
		{
			"it should check pair having sections with full overlaps (sections being equals)",
			fields{
				first:  Section{start: 2, end: 8},
				second: Section{start: 2, end: 8},
			},
			true,
			true,
		},
		{
			"it should check pair not having sections with full overlaps",
			fields{
				first:  Section{start: 2, end: 8},
				second: Section{start: 4, end: 10},
			},
			true,
			false,
		},
		{
			"it should check pair having sections with overlaps",
			fields{
				first:  Section{start: 2, end: 8},
				second: Section{start: 4, end: 10},
			},
			false,
			true,
		},
		{
			"it should check pair having sections with overlaps (example 1)",
			fields{
				first:  Section{start: 5, end: 7},
				second: Section{start: 7, end: 9},
			},
			false,
			true,
		},
		{
			"it should check pair not having sections with overlaps",
			fields{
				first:  Section{start: 5, end: 7},
				second: Section{start: 8, end: 9},
			},
			false,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pair{
				first:  tt.fields.first,
				second: tt.fields.second,
			}
			if got := p.sectionsOverlap(tt.fully); got != tt.want {
				t.Errorf("sectionsOverlap() = %v, want %v", got, tt.want)
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
			args{`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`},
			2,
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
			args{`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`},
			4,
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

func TestSolution1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"result for solution1", 526},
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
		{"result for solution1", 886},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution2(); got != tt.want {
				t.Errorf("Solution2() = %v, want %v", got, tt.want)
			}
		})
	}
}
