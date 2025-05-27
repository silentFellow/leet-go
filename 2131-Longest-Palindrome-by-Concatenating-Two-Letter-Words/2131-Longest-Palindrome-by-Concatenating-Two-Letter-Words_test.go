package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{words: []string{"lc", "cl", "gg"}},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{words: []string{"ab", "ty", "yt", "lc", "cl", "ab"}},
			want: 8,
		},
		{
			name: "Example 3",
			args: args{words: []string{"cc", "ll", "xx"}},
			want: 2,
		},
		{
			name: "No palindrome possible",
			args: args{words: []string{"ab", "cd", "ef"}},
			want: 0,
		},
		{
			name: "Single palindromic word",
			args: args{words: []string{"aa"}},
			want: 2,
		},
		{
			name: "Multiple palindromic words",
			args: args{words: []string{"aa", "bb", "cc"}},
			want: 2,
		},
		{
			name: "Pair and palindromic center",
			args: args{words: []string{"ab", "ba", "cc"}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.words); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
