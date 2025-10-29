package leetcode

import "testing"

func Test_smallestNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"n=1", args{1}, 1},
		{"n=2", args{2}, 3},
		{"n=3", args{3}, 3},
		{"n=5", args{5}, 7},
		{"n=10", args{10}, 15},
		{"n=1000", args{1000}, 1023},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestNumber(tt.args.n); got != tt.want {
				t.Errorf("smallestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
