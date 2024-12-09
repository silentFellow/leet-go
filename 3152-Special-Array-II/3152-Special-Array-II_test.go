package leetcode

import (
	"reflect"
	"testing"
)

func Test_isArraySpecial(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{"Test 1", args{[]int{3, 4, 1, 2, 6}, [][]int{{0, 4}}}, []bool{false}},
		{"Test 2", args{[]int{4, 3, 1, 6}, [][]int{{0, 2}, {2, 3}}}, []bool{false, true}},
		{"Test 2", args{[]int{1, 1}, [][]int{{0, 1}}}, []bool{false}},
		{"Test 4", args{[]int{2, 2}, [][]int{{0, 0}}}, []bool{true}},
		{"Test 5", args{[]int{3, 7, 8}, [][]int{{0, 2}}}, []bool{false}},
		{"Test 6", args{[]int{3, 1, 3, 10, 8, 9}, [][]int{{1, 1}}}, []bool{true}},
		{"Test 7", args{[]int{7, 1, 7, 3}, [][]int{{0, 1}}}, []bool{false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isArraySpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
