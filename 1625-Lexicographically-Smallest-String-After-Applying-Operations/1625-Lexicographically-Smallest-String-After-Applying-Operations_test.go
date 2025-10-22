package leetcode

import "testing"

func Test_findLexSmallestString(t *testing.T) {
	type args struct {
		s string
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example1", args{"5525", 9, 2}, "2050"},
		{"example2", args{"74", 5, 1}, "24"},
		{"example3", args{"0011", 4, 2}, "0011"},
		{"all_zeros", args{"0000", 7, 3}, "0000"},
		{"all_nines", args{"9999", 1, 1}, "0000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLexSmallestString(tt.args.s, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findLexSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
