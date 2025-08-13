package leetcode

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"n=27", args{27}, true},
		{"n=9", args{9}, true},
		{"n=45", args{45}, false},
		{"n=0", args{0}, false},
		{"n=-3", args{-3}, false},
		{"n=1", args{1}, true},
		{"n=1162261467", args{1162261467}, true},
		{"n=243", args{243}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
