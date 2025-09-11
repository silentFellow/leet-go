package leetcode

import "testing"

func Test_sortVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example1", args{"lEetcOde"}, "lEOtcede"},
		{"example2", args{"lYmpH"}, "lYmpH"},
		{"single vowel lowercase", args{"a"}, "a"},
		{"single vowel uppercase", args{"A"}, "A"},
		{"no vowels", args{"bcdfg"}, "bcdfg"},
		{"alternating vowels and consonants", args{"a1e2i3o4u"}, "a1e2i3o4u"},
		{"repeated vowels", args{"eeaaUU"}, "UUaaee"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortVowels(tt.args.s); got != tt.want {
				t.Errorf("sortVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
