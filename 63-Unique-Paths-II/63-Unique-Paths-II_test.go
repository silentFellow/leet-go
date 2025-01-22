package leetcode

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{obstacleGrid: [][]int{{0, 1}, {0, 0}}},
			want: 1,
		},
		{
			name: "No obstacles",
			args: args{obstacleGrid: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
			want: 6,
		},
		{
			name: "Full obstacle in first row",
			args: args{obstacleGrid: [][]int{{1, 1}, {0, 0}}},
			want: 0,
		},
		{
			name: "Full obstacle in first column",
			args: args{obstacleGrid: [][]int{{0, 0}, {1, 0}}},
			want: 1,
		},
		{
			name: "Single cell with obstacle",
			args: args{obstacleGrid: [][]int{{1}}},
			want: 0,
		},
		{
			name: "Single cell without obstacle",
			args: args{obstacleGrid: [][]int{{0}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
