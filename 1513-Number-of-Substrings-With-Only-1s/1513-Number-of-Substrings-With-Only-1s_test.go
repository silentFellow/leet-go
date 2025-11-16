package leetcode

import "testing"

func Test_numSub(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{"0110111"}, 9},
		{"example2", args{"101"}, 2},
		{"example3", args{"111111"}, 21},
		{"single_zero", args{"0"}, 0},
		{"single_one", args{"1"}, 1},
		{"alternating", args{"1010101"}, 4},
		{"all_zeros", args{"00000"}, 0},
		{"all_ones", args{"11111"}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSub(tt.args.s); got != tt.want {
				t.Errorf("numSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
