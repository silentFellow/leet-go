package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{"abccccdd"}, 7},
		{"example2", args{"a"}, 1},
		{"all unique", args{"abc"}, 1},
		{"all pairs", args{"aabbcc"}, 6},
		{"mixed case", args{"Aa"}, 1},
		{"long odd", args{"aaaabbbbccc"}, 11},
		{"empty", args{""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
