package leetcode

import "testing"

func Test_maximumLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"all even", args{[]int{2, 4, 6, 8}}, 4},
		{"all odd", args{[]int{1, 3, 5, 7}}, 4},
		{"alternating even odd", args{[]int{1, 2, 1, 2, 1, 2}}, 6},
		{"alternating odd even", args{[]int{2, 1, 2, 1, 2, 1}}, 6},
		{"mixed", args{[]int{1, 2, 3, 4}}, 4},
		{"short", args{[]int{1, 3}}, 2},
		{"complex", args{[]int{1, 2, 1, 1, 2, 1, 2}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.nums); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
