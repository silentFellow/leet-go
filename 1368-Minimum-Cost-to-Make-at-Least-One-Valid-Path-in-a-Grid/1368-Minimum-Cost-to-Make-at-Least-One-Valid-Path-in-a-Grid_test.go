package leetcode

import (
	"testing"
)

func Test_minCost(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{1, 1, 1, 1},
				{2, 2, 2, 2},
				{1, 1, 1, 1},
				{2, 2, 2, 2},
			},
			want: 3,
		},
		{
			name: "Example 2",
			grid: [][]int{
				{1, 1, 3},
				{3, 2, 2},
				{1, 1, 4},
			},
			want: 0,
		},
		{
			name: "Example 3",
			grid: [][]int{
				{1, 2},
				{4, 3},
			},
			want: 1,
		},
		{
			name: "Single cell",
			grid: [][]int{
				{1},
			},
			want: 0,
		},
		{
			name: "Large grid",
			grid: [][]int{
				{1, 1, 1, 1, 1},
				{2, 2, 2, 2, 2},
				{1, 1, 1, 1, 1},
				{2, 2, 2, 2, 2},
				{1, 1, 1, 1, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.grid); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
