package leetcode

import "testing"

func Test_maxSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{[]int{1, 2}, 1}, 3},
		{"Example 2", args{[]int{-1, -2, -3, -4, -5}, 4}, -10},
		{"Example 3", args{[]int{-5, 1, 2, -3, 4}, 2}, 4},
		{"Single element", args{[]int{7}, 1}, 7},
		{"All positives", args{[]int{2, 4, 6, 8}, 2}, 20},
		{"Mixed, k=3", args{[]int{1, -2, 3, 4, -5, 6}, 3}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
