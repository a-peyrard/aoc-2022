package day13

import (
	"reflect"
	"testing"
)

const InputExample = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

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

func Test_parsePacket(t *testing.T) {
	type args struct {
		rawPacket string
	}
	tests := []struct {
		name string
		args args
		want packet
	}{
		{
			"it should parse an empty packet",
			args{"[]"},
			[]interface{}{},
		},
		{
			"it should parse a packet with numbers",
			args{"[1,2,3]"},
			[]interface{}{1, 2, 3},
		},
		{
			"it should parse a packet with only one nested list",
			args{"[[8,7,6]]"},
			[]interface{}{[]interface{}{8, 7, 6}},
		},
		{
			"it should parse a complex packet",
			args{"[1,[2,[3,[4,[5,6,7]]]],8,9]"},
			[]interface{}{
				1,
				[]interface{}{
					2,
					[]interface{}{
						3,
						[]interface{}{
							4,
							[]interface{}{
								5,
								6,
								7,
							},
						},
					},
				},
				8,
				9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePacket(tt.args.rawPacket); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pairOfPackets_areInOrder(t *testing.T) {
	tests := []struct {
		name string
		p    pairOfPackets
		want bool
	}{
		{
			"it should validate pair 1 from the example",
			[2]packet{
				parsePacket("[1,1,3,1,1]"),
				parsePacket("[1,1,5,1,1]"),
			},
			true,
		},
		{
			"it should validate pair 2 from the example",
			[2]packet{
				parsePacket("[[1],[2,3,4]]"),
				parsePacket("[[1],4]"),
			},
			true,
		},
		{
			"it should validate pair 3 from the example",
			[2]packet{
				parsePacket("[9]"),
				parsePacket("[[8,7,6]]"),
			},
			false,
		},
		{
			"it should validate pair 4 from the example",
			[2]packet{
				parsePacket("[[4,4],4,4]"),
				parsePacket("[[4,4],4,4,4]"),
			},
			true,
		},
		{
			"it should validate pair 5 from the example",
			[2]packet{
				parsePacket("[7,7,7,7]"),
				parsePacket("[7,7,7]"),
			},
			false,
		},
		{
			"it should validate pair 6 from the example",
			[2]packet{
				parsePacket("[]"),
				parsePacket("[3]"),
			},
			true,
		},
		{
			"it should validate pair 7 from the example",
			[2]packet{
				parsePacket("[[1],[2,3,4]]"),
				parsePacket("[[1],4]"),
			},
			true,
		},
		{
			"it should validate pair 8 from the example",
			[2]packet{
				parsePacket("[1,[2,[3,[4,[5,6,7]]]],8,9]"),
				parsePacket("[1,[2,[3,[4,[5,6,0]]]],8,9]"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.areInOrder(); got != tt.want {
				t.Errorf("areInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
