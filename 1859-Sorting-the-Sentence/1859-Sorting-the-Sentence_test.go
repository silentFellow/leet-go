package leetcode

import "testing"

func Test_sortSentence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{s: "is2 sentence4 This1 a3"},
			want: "This is a sentence",
		},
		{
			name: "example2",
			args: args{s: "Myself2 Me1 I4 and3"},
			want: "Me Myself and I",
		},
		{
			name: "single_word",
			args: args{s: "word1"},
			want: "word",
		},
		{
			name: "two_words",
			args: args{s: "second2 first1"},
			want: "first second",
		},
		{
			name: "mixed_case",
			args: args{s: "b2 A1"},
			want: "A b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortSentence(tt.args.s); got != tt.want {
				t.Errorf("sortSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
