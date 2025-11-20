package leetcode

import "testing"

func Test_intersectionSizeTwo(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{"Example 1", [][]int{{1, 3}, {3, 7}, {8, 9}}, 5},
		{"Example 2", [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}, 3},
		{"Example 3", [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}, 5},
		{"Single interval", [][]int{{1, 2}}, 2},
		{"Non-overlapping intervals", [][]int{{1, 2}, {3, 4}, {5, 6}}, 6},
		{"Large interval", [][]int{{0, 100000000}}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectionSizeTwo(tt.intervals); got != tt.want {
				t.Errorf("intersectionSizeTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
