package leetcode

import "testing"

func Test_minZeroArray(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums:    []int{2, 0, 2},
				queries: [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}},
			},
			want: 2,
		},
		{
			name: "example2",
			args: args{
				nums:    []int{4, 3, 2, 1},
				queries: [][]int{{1, 3, 2}, {0, 2, 1}},
			},
			want: -1,
		},
		{
			name: "all zeros",
			args: args{
				nums:    []int{0, 0, 0},
				queries: [][]int{{0, 2, 1}},
			},
			want: 0,
		},
		{
			name: "single element",
			args: args{
				nums:    []int{5},
				queries: [][]int{{0, 0, 5}},
			},
			want: 1,
		},
		{
			name: "no queries",
			args: args{
				nums:    []int{1, 2, 3},
				queries: [][]int{},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minZeroArray(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("minZeroArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
