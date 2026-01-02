package leetcode

import "testing"

func Test_bestClosingTime(t *testing.T) {
	type args struct {
		customers string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"YYNY"}, 2},
		{"Example 2", args{"NNNNN"}, 0},
		{"Example 3", args{"YYYY"}, 4},
		{"Mixed 2", args{"NYYYNNNYNN"}, 4},
		{"Single N", args{"N"}, 0},
		{"Single Y", args{"Y"}, 1},
		{"All Y but last N", args{"YYYYN"}, 4},
		{"All N but last Y", args{"NNNNY"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestClosingTime(tt.args.customers); got != tt.want {
				t.Errorf("bestClosingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
