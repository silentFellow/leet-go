package leetcode

import "testing"

func Test_maximumLength(t *testing.T) {
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
			args: args{s: "aaaa"},
			want: 2,
		},
		{
			name: "example2",
			args: args{s: "abcdef"},
			want: -1,
		},
		{
			name: "example3",
			args: args{s: "abcaba"},
			want: 1,
		},
		{
			name: "test1",
			args: args{s: "aaabbbccc"},
			want: 1,
		},
		{
			name: "test2",
			args: args{s: "zzzzzz"},
			want: 4,
		},
		{
			name: "test3",
			args: args{s: "aabbccddeeffgg"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.s); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
