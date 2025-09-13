package leetcode

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"n=0", args{0}, 0},
		{"n=3", args{3}, 0},
		{"n=5", args{5}, 1},
		{"n=10", args{10}, 2},
		{"n=25", args{25}, 6},
		{"n=100", args{100}, 24},
		{"n=125", args{125}, 31},
		{"n=1000", args{1000}, 249},
		{"n=10000", args{10000}, 2499},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
