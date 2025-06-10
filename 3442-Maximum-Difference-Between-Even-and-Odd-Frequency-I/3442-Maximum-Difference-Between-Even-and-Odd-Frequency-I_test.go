package leetcode

import "testing"

func Test_maxDifference(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{s: "aaaaabbc"},
			want: 3,
		},
		{
			name: "example 2",
			args: args{s: "abcabcab"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDifference(tt.args.s); got != tt.want {
				t.Errorf("maxDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
