package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{s: "babad"},
			want: "bab", // "aba" is also a valid answer
		},
		{
			name: "Example 2",
			args: args{s: "cbbd"},
			want: "bb",
		},
		{
			name: "Single character",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "Empty string",
			args: args{s: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
