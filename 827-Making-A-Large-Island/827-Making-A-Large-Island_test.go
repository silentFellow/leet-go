package leetcode

import "testing"

func Test_largestIsland(t *testing.T) {
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
			args: args{grid: [][]int{{1, 0}, {0, 1}}},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{grid: [][]int{{1, 1}, {1, 0}}},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{grid: [][]int{{1, 1}, {1, 1}}},
			want: 4,
		},
		{
			name: "Single element 0",
			args: args{grid: [][]int{{0}}},
			want: 1,
		},
		{
			name: "Single element 1",
			args: args{grid: [][]int{{1}}},
			want: 1,
		},
		{
			name: "All zeros",
			args: args{grid: [][]int{{0, 0}, {0, 0}}},
			want: 1,
		},
		{
			name: "All ones",
			args: args{grid: [][]int{{1, 1}, {1, 1}}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestIsland(tt.args.grid); got != tt.want {
				t.Errorf("largestIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
