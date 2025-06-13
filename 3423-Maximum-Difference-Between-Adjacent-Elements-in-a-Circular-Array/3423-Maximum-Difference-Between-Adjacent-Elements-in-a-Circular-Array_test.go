package leetcode

import "testing"

func Test_maxAdjacentDistance(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{1, 2, 4}}, 3},
		{"example2", args{[]int{-5, -10, -5}}, 5},
		{"two elements", args{[]int{7, 2}}, 5},
		{"all same", args{[]int{3, 3, 3}}, 0},
		{"negatives", args{[]int{-100, 100}}, 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAdjacentDistance(tt.args.nums); got != tt.want {
				t.Errorf("maxAdjacentDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
