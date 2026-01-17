package leetcode

import "testing"

func Test_largestSquareArea(t *testing.T) {
	type args struct {
		bottomLeft [][]int
		topRight   [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				bottomLeft: [][]int{{1, 1}, {2, 2}, {3, 1}},
				topRight:   [][]int{{3, 3}, {4, 4}, {6, 6}},
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				bottomLeft: [][]int{{1, 1}, {1, 3}, {1, 5}},
				topRight:   [][]int{{5, 5}, {5, 7}, {5, 9}},
			},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{
				bottomLeft: [][]int{{1, 1}, {2, 2}, {1, 2}},
				topRight:   [][]int{{3, 3}, {4, 4}, {3, 4}},
			},
			want: 1,
		},
		{
			name: "Example 4",
			args: args{
				bottomLeft: [][]int{{1, 1}, {3, 3}, {3, 1}},
				topRight:   [][]int{{2, 2}, {4, 4}, {4, 2}},
			},
			want: 0,
		},
		{
			name: "No overlap",
			args: args{
				bottomLeft: [][]int{{1, 1}, {10, 10}},
				topRight:   [][]int{{2, 2}, {11, 11}},
			},
			want: 0,
		},
		{
			name: "Full overlap",
			args: args{
				bottomLeft: [][]int{{1, 1}, {1, 1}},
				topRight:   [][]int{{5, 5}, {5, 5}},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSquareArea(tt.args.bottomLeft, tt.args.topRight); got != tt.want {
				t.Errorf("largestSquareArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
