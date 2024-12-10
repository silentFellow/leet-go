package leetcode

import "testing"

func Test_isArraySpecial(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test 1", args{[]int{1}}, true},
		{"Test 1", args{[]int{2, 1, 4}}, true},
		{"Test 1", args{[]int{4, 3, 1, 6}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial(tt.args.nums); got != tt.want {
				t.Errorf("isArraySpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
