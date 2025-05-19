package leetcode

import "testing"

func Test_triangleType(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"equilateral", args{[]int{3, 3, 3}}, "equilateral"},
		{"scalene", args{[]int{3, 4, 5}}, "scalene"},
		{"isosceles 1", args{[]int{5, 5, 8}}, "isosceles"},
		{"isosceles 2", args{[]int{7, 10, 7}}, "isosceles"},
		{"none 1", args{[]int{1, 2, 3}}, "none"},
		{"none 2", args{[]int{10, 1, 1}}, "none"},
		{"isosceles 3", args{[]int{2, 3, 2}}, "isosceles"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleType(tt.args.nums); got != tt.want {
				t.Errorf("triangleType() = %v, want %v", got, tt.want)
			}
		})
	}
}
