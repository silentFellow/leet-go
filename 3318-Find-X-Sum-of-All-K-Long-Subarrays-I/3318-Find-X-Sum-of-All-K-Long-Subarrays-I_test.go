package leetcode

import (
	"reflect"
	"testing"
)

func Test_findXSum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		x    int
		want []int
	}{
		{[]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2, []int{6, 10, 12}},
		{[]int{3, 8, 7, 8, 7, 5}, 2, 2, []int{11, 15, 15, 15, 12}},
		{[]int{1, 2, 3, 4, 5}, 3, 1, []int{3, 4, 5}},
		{[]int{5, 5, 5, 5, 5}, 3, 2, []int{15, 15, 15}},
		{[]int{1}, 1, 1, []int{1}},
		{[]int{1, 2, 3}, 3, 3, []int{6}},
	}
	for _, tt := range tests {
		got := findXSum(tt.nums, tt.k, tt.x)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findXSum(%v, %d, %d) = %v, want %v", tt.nums, tt.k, tt.x, got, tt.want)
		}
	}
}
