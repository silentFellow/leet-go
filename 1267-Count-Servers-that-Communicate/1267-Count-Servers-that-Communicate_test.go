package leetcode

import "testing"

func Test_countServers(t *testing.T) {
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
			args: args{grid: [][]int{{1, 0}, {0, 1}}},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{grid: [][]int{{1, 0}, {1, 1}}},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{grid: [][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}},
			want: 4,
		},
		{
			name: "Single server",
			args: args{grid: [][]int{{1}}},
			want: 0,
		},
		{
			name: "All servers communicate",
			args: args{grid: [][]int{{1, 1}, {1, 1}}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countServers(tt.args.grid); got != tt.want {
				t.Errorf("countServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
