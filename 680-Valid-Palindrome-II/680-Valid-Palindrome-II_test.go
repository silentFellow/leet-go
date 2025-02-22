package leetcode

import "testing"

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty string",
			args: args{s: ""},
			want: true,
		},
		{
			name: "Single character",
			args: args{s: "a"},
			want: true,
		},
		{
			name: "Palindrome without deletion",
			args: args{s: "aba"},
			want: true,
		},
		{
			name: "Palindrome with one deletion",
			args: args{s: "abca"},
			want: true,
		},
		{
			name: "Not a palindrome",
			args: args{s: "abc"},
			want: false,
		},
		{
			name: "Long palindrome with one deletion",
			args: args{s: "deeee"},
			want: true,
		},
		{
			name: "Long non-palindrome",
			args: args{s: "abcdef"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
