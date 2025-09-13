package leetcode

import "testing"

func Test_doesAliceWin(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1 - Alice wins",
			args: args{s: "leetcoder"},
			want: true,
		},
		{
			name: "Example 2 - Bob wins",
			args: args{s: "bbcd"},
			want: false,
		},
		{
			name: "All vowels, odd count - Alice wins",
			args: args{s: "aeiou"},
			want: true,
		},
		{
			name: "All vowels, even count - Bob wins",
			args: args{s: "aeio"},
			want: false,
		},
		{
			name: "No vowels - Bob wins",
			args: args{s: "bcdfg"},
			want: false,
		},
		{
			name: "Single vowel - Alice wins",
			args: args{s: "a"},
			want: true,
		},
		{
			name: "Single consonant - Bob wins",
			args: args{s: "b"},
			want: false,
		},
		{
			name: "Mixed, even vowels - Bob wins",
			args: args{s: "abcde"},
			want: false,
		},
		{
			name: "Mixed, odd vowels - Alice wins",
			args: args{s: "abcdei"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doesAliceWin(tt.args.s); got != tt.want {
				t.Errorf("doesAliceWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
