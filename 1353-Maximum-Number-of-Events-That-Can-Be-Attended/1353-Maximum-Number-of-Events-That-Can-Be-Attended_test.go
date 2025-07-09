package leetcode

import "testing"

func Test_maxEvents(t *testing.T) {
	type args struct {
		events [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[][]int{{1, 2}, {2, 3}, {3, 4}}}, 3},
		{"Example 2", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}}, 4},
		{"Max Overlap", args{[][]int{{1, 10}, {1, 10}, {1, 10}, {1, 10}}}, 4},
		{"Large Range", args{[][]int{{1, 100}, {50, 100}, {60, 70}, {90, 95}}}, 4},
		{"Single Day Events", args{[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}}, 4},
		{"Same End Day", args{[][]int{{1, 2}, {2, 2}, {1, 2}, {1, 2}}}, 2},
		{"No Overlap", args{[][]int{{1, 2}, {3, 4}, {5, 6}}}, 3},
		{"Disjoint Intervals", args{[][]int{{1, 2}, {4, 5}, {7, 8}, {10, 11}}}, 4},
		{"All Same Day", args{[][]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}}}, 1},
		{"Early Start Late End", args{[][]int{{1, 3}, {2, 4}, {3, 5}, {4, 6}, {5, 7}}}, 5},
		{"Completely Overlap", args{[][]int{{1, 3}, {1, 3}, {1, 3}, {1, 3}}}, 3},
		{"Random Intervals", args{[][]int{{1, 3}, {5, 8}, {6, 10}, {9, 11}, {12, 15}}}, 5},
		{"Edge Case", args{[][]int{{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEvents(tt.args.events); got != tt.want {
				t.Errorf("maxEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
