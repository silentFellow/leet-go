package leetcode

import "testing"

func Test_repeatedStringMatch(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{a: "abcd", b: "cdabcdab"},
			want: 3,
		},
		{
			name: "example2",
			args: args{a: "a", b: "aa"},
			want: 2,
		},
		{
			name: "b is already a substring of a",
			args: args{a: "abc", b: "bc"},
			want: 1,
		},
		{
			name: "b is longer than a but not a substring",
			args: args{a: "abc", b: "def"},
			want: -1,
		},
		{
			name: "a is empty",
			args: args{a: "", b: "abc"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedStringMatch(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("repeatedStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
