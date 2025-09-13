package leetcode

import "testing"

func Test_maxFreqSum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{"successes"}, 6},
		{"example2", args{"aeiaeia"}, 3},
		{"no vowels", args{"bcdfg"}, 1},
		{"no consonants", args{"aeiou"}, 1},
		{"all same letter vowel", args{"aaaaa"}, 5},
		{"all same letter consonant", args{"zzzzz"}, 5},
		{"single char vowel", args{"a"}, 1},
		{"single char consonant", args{"z"}, 1},
		{"empty string", args{""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFreqSum(tt.args.s); got != tt.want {
				t.Errorf("maxFreqSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
