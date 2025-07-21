package leetcode

import "testing"

func Test_makeFancyString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example1", args{"leeetcode"}, "leetcode"},
		{"example2", args{"aaabaaaa"}, "aabaa"},
		{"example3", args{"aab"}, "aab"},
		{"single_char", args{"a"}, "a"},
		{"all_same", args{"aaaaaa"}, "aa"},
		{"no_triple", args{"abcde"}, "abcde"},
		{"long_repeats", args{"bbbbbbbbb"}, "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeFancyString(tt.args.s); got != tt.want {
				t.Errorf("makeFancyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
