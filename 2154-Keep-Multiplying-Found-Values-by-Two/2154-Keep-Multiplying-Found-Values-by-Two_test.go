package leetcode

import "testing"

func Test_findFinalValue(t *testing.T) {
	type args struct {
		nums     []int
		original int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{5, 3, 6, 1, 12}, 3}, 24},
		{"Example 2", args{[]int{2, 7, 9}, 4}, 4},
		{"No match", args{[]int{1, 2, 3}, 8}, 8},
		{"Multiple doubles", args{[]int{2, 4, 8, 16}, 2}, 32},
		{"Single element match", args{[]int{10}, 10}, 20},
		{"Single element no match", args{[]int{10}, 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFinalValue(tt.args.nums, tt.args.original); got != tt.want {
				t.Errorf("findFinalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
