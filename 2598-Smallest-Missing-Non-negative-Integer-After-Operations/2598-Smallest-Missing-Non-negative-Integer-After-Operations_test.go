package leetcode

import "testing"

func Test_findSmallestInteger(t *testing.T) {
	type args struct {
		nums  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example1", args{[]int{1, -10, 7, 13, 6, 8}, 7}, 2},
		{"AllZero", args{[]int{0, 0, 0, 0}, 1}, 4},
		{"Sequential", args{[]int{0, 1, 2, 3, 4}, 5}, 5},
		{"NegativeNumbers", args{[]int{-1, -2, -3}, 2}, 2},
		{"LargeValue", args{[]int{100, 200, 300}, 100}, 1},
		{"Duplicates", args{[]int{3, 0, 3, 2, 4, 2, 1, 1, 0, 4}, 5}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSmallestInteger(tt.args.nums, tt.args.value); got != tt.want {
				t.Errorf("findSmallestInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
