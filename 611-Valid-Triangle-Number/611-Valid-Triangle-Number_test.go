package leetcode

import "testing"

func Test_triangleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{2, 2, 3, 4}}, 3},
		{"example2", args{[]int{4, 2, 3, 4}}, 4},
		{"single element", args{[]int{1}}, 0},
		{"no triangle", args{[]int{1, 1, 3, 4}}, 0},
		{"all zeros", args{[]int{0, 0, 0}}, 0},
		{"one valid triangle", args{[]int{7, 1, 1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleNumber(tt.args.nums); got != tt.want {
				t.Errorf("triangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
