package leetcode

import "testing"

func Test_canBeTypedWords(t *testing.T) {
	type args struct {
		text          string
		brokenLetters string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{text: "hello world", brokenLetters: "ad"},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{text: "leet code", brokenLetters: "lt"},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{text: "leet code", brokenLetters: "e"},
			want: 0,
		},
		{
			name: "No broken letters",
			args: args{text: "a b c", brokenLetters: ""},
			want: 3,
		},
		{
			name: "All letters broken",
			args: args{text: "a b c", brokenLetters: "abc"},
			want: 0,
		},
		{
			name: "Single word, no broken",
			args: args{text: "hello", brokenLetters: ""},
			want: 1,
		},
		{
			name: "Single word, all broken",
			args: args{text: "hello", brokenLetters: "helo"},
			want: 0,
		},
		{
			name: "Mixed words",
			args: args{text: "apple banana cherry", brokenLetters: "b"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeTypedWords(tt.args.text, tt.args.brokenLetters); got != tt.want {
				t.Errorf("canBeTypedWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
