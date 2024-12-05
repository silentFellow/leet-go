package leetcode

import "testing"

func Test_numMatchingSubseq(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				s:     "abcde",
				words: []string{"a", "bb", "acd", "ace"},
			},
			want: 3,
		},
		{
			name: "example2",
			args: args{
				s:     "dsahjpjauf",
				words: []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"},
			},
			want: 2,
		},
		{
			name: "no matches",
			args: args{
				s:     "abc",
				words: []string{"d", "e", "f"},
			},
			want: 0,
		},
		{
			name: "all matches",
			args: args{
				s:     "abc",
				words: []string{"a", "b", "c", "ab", "bc", "abc"},
			},
			want: 6,
		},
		{
			name: "empty string s",
			args: args{
				s:     "",
				words: []string{"a", "b", "c"},
			},
			want: 0,
		},
		{
			name: "empty words",
			args: args{
				s:     "abc",
				words: []string{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMatchingSubseq(tt.args.s, tt.args.words); got != tt.want {
				t.Errorf("numMatchingSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSubString(t *testing.T) {
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
			args: args{
				s: "abcde",
				t: "ace",
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				s: "abcde",
				t: "aec",
			},
			want: false,
		},
		{
			name: "empty t",
			args: args{
				s: "abcde",
				t: "",
			},
			want: true,
		},
		{
			name: "empty s",
			args: args{
				s: "",
				t: "a",
			},
			want: false,
		},
		{
			name: "both empty",
			args: args{
				s: "",
				t: "",
			},
			want: true,
		},
		{
			name: "single character match",
			args: args{
				s: "a",
				t: "a",
			},
			want: true,
		},
		{
			name: "single character no match",
			args: args{
				s: "a",
				t: "b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubString(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}
