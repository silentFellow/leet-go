package leetcode

import "testing"

func Test_reorderedPowerOf2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"n=1", args{1}, true},
		{"n=10", args{10}, false},
		{"n=46", args{46}, true},
		{"n=65", args{65}, false},
		{"n=100", args{100}, false},
		{"n=61", args{61}, true},
		{"n=131072", args{131072}, true},
		{"n=368712509", args{368712509}, true},
		{"n=842717231", args{842717231}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderedPowerOf2(tt.args.n); got != tt.want {
				t.Errorf("reorderedPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
