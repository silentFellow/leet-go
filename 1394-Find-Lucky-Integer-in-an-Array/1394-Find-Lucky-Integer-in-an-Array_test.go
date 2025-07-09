package leetcode

import "testing"

func Test_findLucky(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{2, 2, 3, 4}}, 2},
		{"example2", args{[]int{1, 2, 2, 3, 3, 3}}, 3},
		{"example3", args{[]int{2, 2, 2, 3, 3}}, -1},
		{"single_lucky", args{[]int{1, 2, 2}}, 2},
		{"multiple_lucky", args{[]int{2, 1, 2, 3, 4, 3, 5, 4, 3}}, 3},
		{"no_lucky", args{[]int{7, 8, 9}}, -1},
		{"all_same", args{[]int{4, 4, 4, 4}}, 4},
		{"large_lucky", args{[]int{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLucky(tt.args.arr); got != tt.want {
				t.Errorf("findLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
