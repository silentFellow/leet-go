package leetcode

import "testing"

func Test_countOdds(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"low and high are same even", args{2, 2}, 0},
		{"low and high are same odd", args{3, 3}, 1},
		{"range with both odd", args{3, 7}, 3},
		{"range with both even", args{4, 8}, 2},
		{"range with low even, high odd", args{2, 9}, 4},
		{"range with low odd, high even", args{5, 10}, 3},
		{"range with low=0, high=0", args{0, 0}, 0},
		{"range with low=0, high=1", args{0, 1}, 1},
		{"range with low=8, high=10", args{8, 10}, 1},
		{"large range", args{0, 1000000000}, 500000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOdds(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("countOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}
