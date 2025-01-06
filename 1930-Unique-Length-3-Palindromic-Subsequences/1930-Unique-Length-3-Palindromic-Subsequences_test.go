package leetcode

import "testing"

func Test_countPalindromicSubsequence(t *testing.T) {
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
			args: args{s: "aabca"},
			want: 3,
		},
		{
			name: "example2",
			args: args{s: "adc"},
			want: 0,
		},
		{
			name: "example3",
			args: args{s: "bbcbaba"},
			want: 4,
		},
		{
			name: "example4",
			args: args{s: "tlpjzdmtwderpkpmgoyrcxttiheassztncqvnfjeyxxp"},
			want: 161,
		},
		{
			name: "no_palindromes",
			args: args{s: "abcdef"},
			want: 0,
		},
		{
			name: "all_same_characters",
			args: args{s: "aaaaa"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPalindromicSubsequence(tt.args.s); got != tt.want {
				t.Errorf("countPalindromicSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
