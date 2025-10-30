package leetcode

import "testing"

func Test_minNumberOperations(t *testing.T) {
	type args struct {
		target []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"single element", args{[]int{4}}, 4},
		{"all same", args{[]int{1, 1, 1, 1}}, 1},
		{"increasing then decreasing", args{[]int{1, 2, 3, 2, 1}}, 3},
		{"decreasing", args{[]int{5, 4, 3, 2, 1}}, 5},
		{"zigzag", args{[]int{3, 1, 5, 4, 2}}, 7},
		{"example2", args{[]int{3, 1, 1, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOperations(tt.args.target); got != tt.want {
				t.Errorf("minNumberOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
