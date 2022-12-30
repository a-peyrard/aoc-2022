package collection

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		slice *[]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it should sum slices' elements",
			args{&([]int{1, 2, 3})},
			6,
		},
		{
			"it should return 0 for empty slices",
			args{&([]int{})},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.slice); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
