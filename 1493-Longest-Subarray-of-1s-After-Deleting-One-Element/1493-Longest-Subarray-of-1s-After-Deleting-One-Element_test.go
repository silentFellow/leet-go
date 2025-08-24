package leetcode

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{1, 1, 0, 1}}, 3},
		{"Example 2", args{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}}, 5},
		{"Example 3", args{[]int{1, 1, 1}}, 2},
		{"All zeros", args{[]int{0, 0, 0}}, 0},
		{"Single zero", args{[]int{0}}, 0},
		{"Single one", args{[]int{1}}, 0},
		{"Alternating", args{[]int{1, 0, 1, 0, 1}}, 2},
		{"Two zeros together", args{[]int{1, 1, 0, 0, 1, 1}}, 2},
		{"Long run of ones with zero at end", args{[]int{1, 1, 1, 1, 0}}, 4},
		{"Long run of ones with zero at start", args{[]int{0, 1, 1, 1, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
