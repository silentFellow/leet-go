package leetcode

import (
	"reflect"
	"testing"
)

func Test_productQueries(t *testing.T) {
	type args struct {
		n       int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				n:       15,
				queries: [][]int{{0, 1}, {2, 2}, {0, 3}},
			},
			want: []int{2, 4, 64},
		},
		{
			name: "Example 2",
			args: args{
				n:       2,
				queries: [][]int{{0, 0}},
			},
			want: []int{2},
		},
		{
			name: "Single bit",
			args: args{
				n:       8,
				queries: [][]int{{0, 0}},
			},
			want: []int{8},
		},
		{
			name: "All bits",
			args: args{
				n:       7,
				queries: [][]int{{0, 2}},
			},
			want: []int{1 * 2 * 4},
		},
		{
			name: "Large n, single query",
			args: args{
				n:       1023,
				queries: [][]int{{0, 9}},
			},
			want: []int{1 << 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productQueries(tt.args.n, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
