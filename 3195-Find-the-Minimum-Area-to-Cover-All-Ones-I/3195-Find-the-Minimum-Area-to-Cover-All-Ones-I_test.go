package leetcode

import "testing"

func Test_minimumArea(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{grid: [][]int{{0, 1, 0}, {1, 0, 1}}},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{grid: [][]int{{1, 0}, {0, 0}}},
			want: 1,
		},
		{
			name: "All ones",
			args: args{grid: [][]int{{1, 1}, {1, 1}}},
			want: 4,
		},
		{
			name: "Single one",
			args: args{grid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
			want: 1,
		},
		{
			name: "Ones in corners",
			args: args{grid: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumArea(tt.args.grid); got != tt.want {
				t.Errorf("minimumArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
