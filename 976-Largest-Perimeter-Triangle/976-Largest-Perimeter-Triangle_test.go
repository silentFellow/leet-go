package leetcode

import "testing"

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{2, 1, 2}}, 5},
		{"example2", args{[]int{1, 2, 1, 10}}, 0},
		{"all equal", args{[]int{3, 3, 3}}, 9},
		{"no triangle", args{[]int{1, 1, 2}}, 0},
		{"large triangle", args{[]int{3, 6, 2, 3}}, 8},
		{"descending", args{[]int{10, 9, 8, 7}}, 27},
		{"minimum input", args{[]int{1, 1, 1}}, 3},
		{"big numbers", args{[]int{1000000, 1000000, 1000000}}, 3000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPerimeter(tt.args.nums); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
