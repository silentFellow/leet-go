package leetcode

import "testing"

func Test_specialTriplets(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{6, 3, 6}, 1},
		{[]int{0, 1, 0, 0}, 1},
		{[]int{8, 4, 2, 8, 4}, 2},
		{[]int{1, 2, 3}, 0},
		{[]int{0, 0, 0, 0}, 4}, // C(4,3) = 4
	}
	for _, tt := range tests {
		got := specialTriplets(tt.nums)
		if got != tt.want {
			t.Errorf("specialTriplets(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
