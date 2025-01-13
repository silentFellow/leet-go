package leetcode

import "testing"

func Test_minimumLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{s: "abaacbcbb"},
			want: 5,
		},
		{
			name: "example2",
			args: args{s: "aa"},
			want: 2,
		},
		{
			name: "example3",
			args: args{s: "abcabc"},
			want: 6,
		},
		{
			name: "example4",
			args: args{s: "aabbcc"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLength(tt.args.s); got != tt.want {
				t.Errorf("minimumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
