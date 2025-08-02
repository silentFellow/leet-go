package leetcode

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		basket1 []int
		basket2 []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				basket1: []int{4, 2, 2, 2},
				basket2: []int{1, 4, 1, 2},
			},
			want: 1,
		},
		{
			name: "Example 2 - Impossible",
			args: args{
				basket1: []int{2, 3, 4, 1},
				basket2: []int{3, 2, 5, 1},
			},
			want: -1,
		},
		{
			name: "Multiple swaps required",
			args: args{
				basket1: []int{1, 2, 2, 4},
				basket2: []int{2, 3, 3, 1},
			},
			want: -1,
		},
		{
			name: "Already equal",
			args: args{
				basket1: []int{1, 2, 3},
				basket2: []int{1, 2, 3},
			},
			want: 0,
		},
		{
			name: "Uneven total frequency",
			args: args{
				basket1: []int{1, 2},
				basket2: []int{1, 2, 2},
			},
			want: -1,
		},
		{
			name: "Large numbers, but balance possible",
			args: args{
				basket1: []int{1000000000, 1},
				basket2: []int{1000000000, 1},
			},
			want: 0,
		},
		{
			name: "Use minimum element for cheaper swap",
			args: args{
				basket1: []int{84, 80, 43, 8, 80, 88, 43, 14, 100, 88},
				basket2: []int{32, 32, 42, 68, 68, 100, 42, 84, 14, 8},
			},
			want: 48,
		},
		{
			name: "Single item equal",
			args: args{
				basket1: []int{1},
				basket2: []int{1},
			},
			want: 0,
		},
		{
			name: "Single item mismatch",
			args: args{
				basket1: []int{1},
				basket2: []int{2},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.basket1, tt.args.basket2); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
