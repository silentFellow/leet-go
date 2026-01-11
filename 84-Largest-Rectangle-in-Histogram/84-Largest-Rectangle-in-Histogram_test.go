package leetcode

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{heights: []int{2, 1, 5, 6, 2, 3}},
			want: 10,
		},
		{
			name: "Example 2",
			args: args{heights: []int{2, 4}},
			want: 4,
		},
		{
			name: "All same height",
			args: args{heights: []int{3, 3, 3, 3}},
			want: 12,
		},
		{
			name: "Single bar",
			args: args{heights: []int{7}},
			want: 7,
		},
		{
			name: "Decreasing heights",
			args: args{heights: []int{5, 4, 3, 2, 1}},
			want: 9,
		},
		{
			name: "Increasing heights",
			args: args{heights: []int{1, 2, 3, 4, 5}},
			want: 9,
		},
		{
			name: "Zeros in heights",
			args: args{heights: []int{0, 2, 0, 3, 0}},
			want: 3,
		},
		{
			name: "Multiple peaks",
			args: args{heights: []int{2, 1, 2}},
			want: 3,
		},
		{
			name: "Empty input",
			args: args{heights: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
