package leetcode

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{s: "Hello World"},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{s: "   fly me   to   the moon  "},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{s: "luffy is still joyboy"},
			want: 6,
		},
		{
			name: "Single word",
			args: args{s: "hello"},
			want: 5,
		},
		{
			name: "Trailing spaces",
			args: args{s: "hello   "},
			want: 5,
		},
		{
			name: "Empty string",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "Spaces only",
			args: args{s: "     "},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
