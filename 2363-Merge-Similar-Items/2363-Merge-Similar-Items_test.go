package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeSimilarItems(t *testing.T) {
	type args struct {
		items1 [][]int
		items2 [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example1",
			args: args{
				items1: [][]int{{1, 1}, {4, 5}, {3, 8}},
				items2: [][]int{{3, 1}, {1, 5}},
			},
			want: [][]int{{1, 6}, {3, 9}, {4, 5}},
		},
		{
			name: "example2",
			args: args{
				items1: [][]int{{1, 1}, {3, 2}, {2, 3}},
				items2: [][]int{{2, 1}, {3, 2}, {1, 3}},
			},
			want: [][]int{{1, 4}, {2, 4}, {3, 4}},
		},
		{
			name: "example3",
			args: args{
				items1: [][]int{{1, 3}, {2, 2}},
				items2: [][]int{{7, 1}, {2, 2}, {1, 4}},
			},
			want: [][]int{{1, 7}, {2, 4}, {7, 1}},
		},
		{
			name: "no overlap",
			args: args{
				items1: [][]int{{1, 2}, {2, 3}},
				items2: [][]int{{3, 4}, {4, 5}},
			},
			want: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
		},
		{
			name: "empty items1",
			args: args{
				items1: [][]int{},
				items2: [][]int{{1, 1}, {2, 2}},
			},
			want: [][]int{{1, 1}, {2, 2}},
		},
		{
			name: "empty items2",
			args: args{
				items1: [][]int{{1, 1}, {2, 2}},
				items2: [][]int{},
			},
			want: [][]int{{1, 1}, {2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSimilarItems(tt.args.items1, tt.args.items2); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("mergeSimilarItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
