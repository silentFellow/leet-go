package leetcode

import "testing"

func Test_countFairPairs(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example1",
			args: args{
				nums:  []int{0, 1, 7, 4, 4, 5},
				lower: 3,
				upper: 6,
			},
			want: 6,
		},
		{
			name: "example2",
			args: args{
				nums:  []int{1, 7, 9, 2, 5},
				lower: 11,
				upper: 11,
			},
			want: 1,
		},
		{
			name: "no_pairs",
			args: args{
				nums:  []int{1, 2, 3},
				lower: 10,
				upper: 15,
			},
			want: 0,
		},
		{
			name: "all_pairs",
			args: args{
				nums:  []int{1, 1, 1, 1},
				lower: 2,
				upper: 2,
			},
			want: 6,
		},
		{
			name: "large_range",
			args: args{
				nums:  []int{-10, -5, 0, 5, 10},
				lower: -15,
				upper: 15,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countFairPairs(tt.args.nums, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("countFairPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
