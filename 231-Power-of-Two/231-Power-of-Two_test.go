package leetcode

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"n=1", args{1}, true},
		{"n=2", args{2}, true},
		{"n=3", args{3}, false},
		{"n=4", args{4}, true},
		{"n=16", args{16}, true},
		{"n=0", args{0}, false},
		{"n=-2", args{-2}, false},
		{"n=1024", args{1024}, true},
		{"n=218", args{218}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
