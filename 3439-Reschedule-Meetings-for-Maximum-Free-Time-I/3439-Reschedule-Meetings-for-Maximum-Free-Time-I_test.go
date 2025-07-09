package leetcode

import "testing"

func Test_maxFreeTime(t *testing.T) {
	type args struct {
		eventTime int
		k         int
		startTime []int
		endTime   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				eventTime: 5,
				k:         1,
				startTime: []int{1, 3},
				endTime:   []int{2, 5},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				eventTime: 10,
				k:         1,
				startTime: []int{0, 2, 9},
				endTime:   []int{1, 4, 10},
			},
			want: 6,
		},
		{
			name: "Example 3",
			args: args{
				eventTime: 5,
				k:         2,
				startTime: []int{0, 1, 2, 3, 4},
				endTime:   []int{1, 2, 3, 4, 5},
			},
			want: 0,
		},
		{
			name: "Edge: all meetings at start",
			args: args{
				eventTime: 10,
				k:         2,
				startTime: []int{0, 1, 2},
				endTime:   []int{1, 2, 3},
			},
			want: 7,
		},
		{
			name: "Edge: all meetings at end",
			args: args{
				eventTime: 10,
				k:         2,
				startTime: []int{7, 8, 9},
				endTime:   []int{8, 9, 10},
			},
			want: 7,
		},
		{
			name: "Custom: minimal free time",
			args: args{
				eventTime: 5,
				k:         1,
				startTime: []int{0, 2, 4},
				endTime:   []int{1, 3, 5},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFreeTime(tt.args.eventTime, tt.args.k, tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("maxFreeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
