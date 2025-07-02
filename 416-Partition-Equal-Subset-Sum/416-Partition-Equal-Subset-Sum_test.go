package leetcode

import "testing"

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{[]int{1, 5, 11, 5}}, true},
		{"example2", args{[]int{1, 2, 3, 5}}, false},
		{"all equal", args{[]int{1, 1}}, true},
		{"odd sum", args{[]int{1, 2, 5}}, false},
		{"large true", args{[]int{23, 13, 11, 7, 6, 5, 5}}, true},
		{"large false", args{[]int{9, 10, 15, 3, 9, 2, 9, 10, 13, 1}}, false},
		{"single", args{[]int{2}}, false},
		{"all same", args{[]int{4, 4, 4, 4}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
