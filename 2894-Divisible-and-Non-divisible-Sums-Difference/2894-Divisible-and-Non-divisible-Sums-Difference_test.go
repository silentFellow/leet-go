package leetcode

import "testing"

func Test_differenceOfSums(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{10, 3}, 19},
		{"example2", args{5, 6}, 15},
		{"example3", args{5, 1}, -15},
		{"m_greater_than_n", args{3, 10}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceOfSums(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("differenceOfSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
