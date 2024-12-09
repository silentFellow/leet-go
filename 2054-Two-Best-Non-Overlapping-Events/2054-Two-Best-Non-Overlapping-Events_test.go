package leetcode

import "testing"

func Test_maxTwoEvents(t *testing.T) {
	type args struct {
		events [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{events: [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}},
			want: 4,
		},
		{
			name: "example2",
			args: args{events: [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}},
			want: 5,
		},
		{
			name: "example3",
			args: args{events: [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}},
			want: 8,
		},
		{
			name: "no_overlap",
			args: args{events: [][]int{{1, 2, 4}, {3, 4, 3}, {5, 6, 2}}},
			want: 7,
		},
		{
			name: "all_overlap",
			args: args{events: [][]int{{1, 5, 5}, {2, 6, 6}, {3, 7, 7}}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTwoEvents(tt.args.events); got != tt.want {
				t.Errorf("maxTwoEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
