package leetcode

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"n=16", args{16}, true},
		{"n=5", args{5}, false},
		{"n=1", args{1}, true},
		{"n=4", args{4}, true},
		{"n=0", args{0}, false},
		{"n=-4", args{-4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
