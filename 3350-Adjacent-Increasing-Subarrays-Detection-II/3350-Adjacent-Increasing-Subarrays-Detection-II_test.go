package leetcode

import "testing"

func Test_maxIncreasingSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"single element", args{[]int{1}}, 0},
		{"all increasing", args{[]int{1, 2, 3, 4}}, 2},
		{"two adjacent increasing", args{[]int{1, 2, 3, 1, 2, 3}}, 3},
		{"with plateau", args{[]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}}, 2},
		{"no increasing", args{[]int{5, 4, 3, 2, 1}}, 1},
		{"mixed", args{[]int{1, 3, 2, 4, 5, 1, 2, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIncreasingSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("maxIncreasingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
