package leetcode

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})

	tests := []struct {
		name  string
		left  int
		right int
		want  int
	}{
		{"sumRange(0, 2)", 0, 2, 1},
		{"sumRange(2, 5)", 2, 5, -1},
		{"sumRange(0, 5)", 0, 5, -3},
		{"sumRange(1, 3)", 1, 3, -2},
		{"sumRange(3, 3)", 3, 3, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numArray.SumRange(tt.left, tt.right)
			if got != tt.want {
				t.Errorf("SumRange(%d, %d) = %d, want %d", tt.left, tt.right, got, tt.want)
			}
		})
	}
}
