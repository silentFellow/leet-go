package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid word", args{"234Adas"}, true},
		{"too short, no vowel", args{"b3"}, false},
		{"invalid char, no consonant", args{"a3$e"}, false},
		{"only digits", args{"1234"}, false},
		{"no vowel", args{"bcdfg3"}, false},
		{"no consonant", args{"aei3"}, false},
		{"invalid char", args{"abc@"}, false},
		{"all valid, mixed case", args{"Ae3b"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.word); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
