package leetcode

import (
	"reflect"
	"testing"
)

func Test_leftmostBuildingQueries(t *testing.T) {
	type args struct {
		heights []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				heights: []int{6, 4, 8, 5, 2, 7},
				queries: [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}},
			},
			want: []int{2, 5, -1, 5, 2},
		},
		{
			name: "example2",
			args: args{
				heights: []int{5, 3, 8, 2, 6, 1, 4, 6},
				queries: [][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}},
			},
			want: []int{7, 6, -1, 4, 6},
		},
		{
			name: "example3",
			args: args{
				heights: []int{3, 4, 1, 2},
				queries: [][]int{
					{0, 0},
					{0, 1},
					{0, 2},
					{0, 3},
					{1, 0},
					{1, 1},
					{1, 2},
					{1, 3},
					{2, 0},
					{2, 1},
					{2, 2},
					{2, 3},
					{3, 0},
					{3, 1},
					{3, 2},
					{3, 3},
				},
			},
			want: []int{0, 1, -1, -1, 1, 1, -1, -1, -1, -1, 2, 3, -1, -1, 3, 3},
		},
		{
			name: "custom_test_case",
			args: args{
				heights: []int{3, 4, 1, 2, 5},
				queries: [][]int{
					{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
					{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4},
					{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4},
					{3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4},
					{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4},
				},
			},
			want: []int{0, 1, 4, 4, 4, 1, 1, 4, 4, 4, 4, 4, 2, 3, 4, 4, 4, 3, 3, 4, 4, 4, 4, 4, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftmostBuildingQueries(tt.args.heights, tt.args.queries); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("leftmostBuildingQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
