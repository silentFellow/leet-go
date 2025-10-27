package leetcode

import "testing"

func Test_nextBeautifulNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{1}, 22},
		{"Example 2", args{1000}, 1333},
		{"Example 3", args{3000}, 3133},
		{"Edge 0", args{0}, 1},
		{"Edge 22", args{22}, 122},
		{"Edge 999999", args{999999}, 1224444},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextBeautifulNumber(tt.args.n); got != tt.want {
				t.Errorf("nextBeautifulNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
