package leetcode

import "testing"

func Test_maximumDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{7, 1, 5, 4}}, 4},
		{"example2", args{[]int{9, 4, 3, 2}}, -1},
		{"example3", args{[]int{1, 5, 2, 10}}, 9},
		{"all decreasing", args{[]int{5, 4, 3, 2, 1}}, -1},
		{"two elements increasing", args{[]int{1, 2}}, 1},
		{"two elements decreasing", args{[]int{2, 1}}, -1},
		{"duplicates", args{[]int{2, 2, 2, 3}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumDifference(tt.args.nums); got != tt.want {
				t.Errorf("maximumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
