package leetcode

import "testing"

func Test_numWaterBottles(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{9, 3}, 13},
		{"example2", args{15, 4}, 19},
		{"min_bottles", args{1, 2}, 1},
		{"exchange_equals_bottles", args{10, 10}, 11},
		{"exchange_greater_than_bottles", args{5, 10}, 5},
		{"max_bottles", args{100, 2}, 199},
		{"max_exchange", args{100, 100}, 101},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWaterBottles(tt.args.numBottles, tt.args.numExchange); got != tt.want {
				t.Errorf("numWaterBottles() = %v, want %v", got, tt.want)
			}
		})
	}
}
