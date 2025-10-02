package leetcode

import "testing"

func Test_maxBottlesDrunk(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{13, 6}, 15},
		{"example2", args{10, 3}, 13},
		{"min_values", args{1, 1}, 2},
		{"exchange_too_high", args{5, 10}, 5},
		{"exchange_just_possible", args{6, 6}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxBottlesDrunk(tt.args.numBottles, tt.args.numExchange); got != tt.want {
				t.Errorf("maxBottlesDrunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
