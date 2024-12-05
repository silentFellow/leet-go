package leetcode

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{s: "abc", t: "ahbgdc"},
			want: true,
		},
		{
			name: "example2",
			args: args{s: "axc", t: "ahbgdc"},
			want: false,
		},
		{
			name: "empty_s",
			args: args{s: "", t: "ahbgdc"},
			want: true,
		},
		{
			name: "empty_t",
			args: args{s: "abc", t: ""},
			want: false,
		},
		{
			name: "both_empty",
			args: args{s: "", t: ""},
			want: true,
		},
		{
			name: "single_char_match",
			args: args{s: "a", t: "a"},
			want: true,
		},
		{
			name: "single_char_no_match",
			args: args{s: "a", t: "b"},
			want: false,
		},
		{
			name: "longer_s",
			args: args{s: "abcdef", t: "abc"},
			want: false,
		},
		{
			name: "subsequence_at_end",
			args: args{s: "fg", t: "abcdefg"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
