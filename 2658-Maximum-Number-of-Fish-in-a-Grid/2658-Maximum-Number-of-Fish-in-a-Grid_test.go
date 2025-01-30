package leetcode

import "testing"

func Test_findMaxFish(t *testing.T) {
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
			args: args{grid: [][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{grid: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}}},
			want: 1,
		},
		{
			name: "No water cells",
			args: args{grid: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
			want: 0,
		},
		{
			name: "Single water cell",
			args: args{grid: [][]int{{0, 0, 0}, {0, 5, 0}, {0, 0, 0}}},
			want: 5,
		},
		{
			name: "Multiple water cells with same fish count",
			args: args{grid: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxFish(tt.args.grid); got != tt.want {
				t.Errorf("findMaxFish() = %v, want %v", got, tt.want)
			}
		})
	}
}
