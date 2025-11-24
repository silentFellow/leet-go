package leetcode

import "testing"

func Test_maxSumDivThree(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example1", []int{3, 6, 5, 1, 8}, 18},
		{"example2", []int{4}, 0},
		{"example3", []int{1, 2, 3, 4, 4}, 12},
		{"all divisible", []int{3, 6, 9}, 18},
		{"none divisible", []int{1, 2, 4, 5}, 12},
		{"single element divisible", []int{6}, 6},
		{"single element not divisible", []int{7}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumDivThree(tt.nums); got != tt.want {
				t.Errorf("maxSumDivThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
