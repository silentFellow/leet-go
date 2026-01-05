package leetcode

import "testing"

func Test_maxMatrixSum(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int64
	}{
		{"all positive", [][]int{{1, 2}, {3, 4}}, 10},
		{"all negative", [][]int{{-1, -2}, {-3, -4}}, 10},
		{"mixed small", [][]int{{1, -1}, {-1, 1}}, 4},
		{"mixed larger", [][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}, 16},
		{"contains zero", [][]int{{0, -1}, {2, 3}}, 6},
		{"all zero", [][]int{{0, 0}, {0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxMatrixSum(tt.matrix); got != tt.want {
				t.Errorf("maxMatrixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
