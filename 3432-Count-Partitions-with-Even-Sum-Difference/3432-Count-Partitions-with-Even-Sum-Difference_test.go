package leetcode

import "testing"

func Test_countPartitions(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{10, 10, 3, 7, 6}}, 4},
		{"example2", args{[]int{1, 2, 2}}, 0},
		{"example3", args{[]int{2, 4, 6, 8}}, 3},
		{"all_even", args{[]int{2, 2}}, 1},
		{"all_odd_even_count", args{[]int{1, 3, 5, 7}}, 3},
		{"odd_count_odd", args{[]int{1, 3, 5}}, 0},
		{"mixed", args{[]int{1, 2, 3, 4, 5}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPartitions(tt.args.nums); got != tt.want {
				t.Errorf("countPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
