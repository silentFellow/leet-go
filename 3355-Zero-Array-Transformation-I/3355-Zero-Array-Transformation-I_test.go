package leetcode

import "testing"

func Test_isZeroArray(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				nums:    []int{1, 0, 1},
				queries: [][]int{{0, 2}},
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				nums:    []int{4, 3, 2, 1},
				queries: [][]int{{1, 3}, {0, 2}},
			},
			want: false,
		},
		{
			name: "All zeros, no queries",
			args: args{
				nums:    []int{0, 0, 0},
				queries: [][]int{},
			},
			want: true,
		},
		{
			name: "Single element, single query",
			args: args{
				nums:    []int{1},
				queries: [][]int{{0, 0}},
			},
			want: true,
		},
		{
			name: "Single element, not enough queries",
			args: args{
				nums:    []int{2},
				queries: [][]int{{0, 0}},
			},
			want: false,
		},
		{
			name: "Multiple queries, impossible",
			args: args{
				nums:    []int{2},
				queries: [][]int{{0, 0}, {0, 0}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isZeroArray(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("isZeroArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
