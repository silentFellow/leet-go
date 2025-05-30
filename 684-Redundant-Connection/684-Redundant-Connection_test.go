package leetcode

import (
	"reflect"
	"testing"
)

func Test_findRedundantConnection(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{edges: [][]int{{1, 2}, {1, 3}, {2, 3}}},
			want: []int{2, 3},
		},
		{
			name: "Example 2",
			args: args{edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}},
			want: []int{1, 4},
		},
		{
			name: "Case 3",
			args: args{edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {2, 5}}},
			want: []int{2, 5},
		},
		{
			name: "Case 4",
			args: args{edges: [][]int{{1, 2}, {2, 3}, {3, 1}}},
			want: []int{3, 1},
		},
		{
			name: "Case 5",
			args: args{edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {4, 5}}},
			want: []int{4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRedundantConnection(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
