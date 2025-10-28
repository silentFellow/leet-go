package leetcode

import "testing"

func Test_countValidSelections(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Example 1", []int{1, 0, 2, 0, 3}, 2},
		{"Example 2", []int{2, 3, 4, 0, 4, 1, 0}, 0},
		{"Single zero", []int{0}, 2},
		{"Multiple zeros, all others zero", []int{0, 0, 0}, 6},
		{"Zeros at ends", []int{0, 1, 0}, 2},
		{"No valid selection", []int{1, 2, 0, 1}, 0},
		{"All non-zero except one zero", []int{1, 2, 3, 0, 4, 5}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countValidSelections(tt.nums); got != tt.want {
				t.Errorf("countValidSelections(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
