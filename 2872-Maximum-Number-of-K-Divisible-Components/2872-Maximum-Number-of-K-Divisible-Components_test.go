package leetcode

import "testing"

func Test_maxKDivisibleComponents(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		values []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				n:      5,
				edges:  [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}},
				values: []int{1, 8, 1, 4, 4},
				k:      6,
			},
			want: 2,
		},
		{
			name: "example2",
			args: args{
				n:      7,
				edges:  [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
				values: []int{3, 0, 6, 1, 5, 2, 1},
				k:      3,
			},
			want: 3,
		},
		{
			name: "single node",
			args: args{
				n:      1,
				edges:  [][]int{},
				values: []int{5},
				k:      5,
			},
			want: 1,
		},
		{
			name: "no valid split",
			args: args{
				n:      4,
				edges:  [][]int{{0, 1}, {1, 2}, {2, 3}},
				values: []int{1, 2, 3, 4},
				k:      10,
			},
			want: 1,
		},
		{
			name: "all nodes divisible",
			args: args{
				n:      3,
				edges:  [][]int{{0, 1}, {1, 2}},
				values: []int{3, 6, 9},
				k:      3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxKDivisibleComponents(tt.args.n, tt.args.edges, tt.args.values, tt.args.k); got != tt.want {
				t.Errorf("maxKDivisibleComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
