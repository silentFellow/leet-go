package leetcode

import "testing"

func Test_countDays(t *testing.T) {
	type args struct {
		days     int
		meetings [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				days:     10,
				meetings: [][]int{{5, 7}, {1, 3}, {9, 10}},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				days:     5,
				meetings: [][]int{{2, 4}, {1, 3}},
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				days:     6,
				meetings: [][]int{{1, 6}},
			},
			want: 0,
		},
		{
			name: "No meetings",
			args: args{
				days:     7,
				meetings: [][]int{},
			},
			want: 7,
		},
		{
			name: "All days overlap",
			args: args{
				days:     10,
				meetings: [][]int{{1, 10}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDays(tt.args.days, tt.args.meetings); got != tt.want {
				t.Errorf("countDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
