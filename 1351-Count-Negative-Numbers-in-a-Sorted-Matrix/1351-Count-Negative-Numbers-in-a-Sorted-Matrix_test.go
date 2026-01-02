package leetcode

import "testing"

func Test_countNegatives(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{grid: [][]int{
				{4, 3, 2, -1},
				{3, 2, 1, -1},
				{1, 1, -1, -2},
				{-1, -1, -2, -3},
			}},
			want: 8,
		},
		{
			name: "Example 2",
			args: args{grid: [][]int{
				{3, 2},
				{1, 0},
			}},
			want: 0,
		},
		{
			name: "All negatives",
			args: args{grid: [][]int{
				{-1, -2},
				{-3, -4},
			}},
			want: 4,
		},
		{
			name: "All positives",
			args: args{grid: [][]int{
				{5, 4},
				{3, 2},
			}},
			want: 0,
		},
		{
			name: "Single row",
			args: args{grid: [][]int{
				{2, 1, 0, -1, -2},
			}},
			want: 2,
		},
		{
			name: "Single column",
			args: args{grid: [][]int{
				{3},
				{2},
				{-1},
				{-2},
			}},
			want: 2,
		},
		{
			name: "Mixed",
			args: args{grid: [][]int{
				{5, 1, 0},
				{-1, -2, -3},
				{-4, -5, -6},
			}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNegatives(tt.args.grid); got != tt.want {
				t.Errorf("countNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}
