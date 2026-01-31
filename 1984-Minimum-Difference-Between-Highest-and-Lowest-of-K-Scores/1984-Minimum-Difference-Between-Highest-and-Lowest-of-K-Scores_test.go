package leetcode

import "testing"

func Test_minimumDifference(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"single element", args{[]int{90}, 1}, 0},
		{"two elements", args{[]int{9, 4}, 2}, 5},
		{"example 2", args{[]int{9, 4, 1, 7}, 2}, 2},
		{"all same", args{[]int{5, 5, 5, 5}, 3}, 0},
		{"sorted input", args{[]int{1, 2, 3, 4, 5}, 3}, 2},
		{"reverse sorted", args{[]int{5, 4, 3, 2, 1}, 3}, 2},
		{"spread out", args{[]int{20, 200, 300, 1000}, 3}, 280},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDifference(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
