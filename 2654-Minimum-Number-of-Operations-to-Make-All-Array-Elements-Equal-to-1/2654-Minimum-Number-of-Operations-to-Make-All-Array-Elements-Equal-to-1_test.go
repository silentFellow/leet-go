package leetcode

import "testing"

func Test_gcd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"gcd(6, 4)", args{6, 4}, 2},
		{"gcd(10, 15)", args{10, 15}, 5},
		{"gcd(7, 3)", args{7, 3}, 1},
		{"gcd(0, 5)", args{0, 5}, 5},
		{"gcd(5, 0)", args{5, 0}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{2, 6, 3, 4}, 4},
		{"example 2", []int{2, 10, 6, 14}, -1},
		{"all ones", []int{1, 1, 1}, 2},
		{"one present", []int{1, 2, 3}, 2},
		{"no possible", []int{4, 8, 16}, -1},
		{"gcd in subarray", []int{3, 6, 5, 10}, 4},
		{"pair gcd 1", []int{2, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.nums); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
