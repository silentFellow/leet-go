package leetcode

import "testing"

func Test_minimumOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example1", []int{1, 2, 3, 4}, 3},
		{"example2", []int{3, 6, 9}, 0},
		{"all need ops", []int{1, 2, 4, 5}, 4},
		{"none need ops", []int{3, 6, 9, 12}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.nums); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
