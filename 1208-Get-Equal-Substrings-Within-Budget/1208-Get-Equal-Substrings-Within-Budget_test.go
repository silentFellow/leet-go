package leetcode

import "testing"

func Test_equalSubstring(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{s: "abcd", t: "bcdf", maxCost: 3},
			want: 3,
		},
		{
			name: "example2",
			args: args{s: "abcd", t: "cdef", maxCost: 3},
			want: 1,
		},
		{
			name: "example3",
			args: args{s: "abcd", t: "acde", maxCost: 0},
			want: 1,
		},
		{
			name: "example4",
			args: args{s: "krrgw", t: "zjxss", maxCost: 18},
			want: 2,
		},
		{
			name: "no change needed",
			args: args{s: "abcd", t: "abcd", maxCost: 10},
			want: 4,
		},
		{
			name: "max cost exceeded",
			args: args{s: "abcd", t: "wxyz", maxCost: 5},
			want: 0,
		},
		{
			name: "partial change",
			args: args{s: "abcd", t: "abcf", maxCost: 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSubstring(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
