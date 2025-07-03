package leetcode

import "testing"

func Test_freqAlphabets(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example1", args{"10#11#12"}, "jkab"},
		// {"example2", args{"1326#"}, "acz"},
		// {"single_digits", args{"123456789"}, "abcdefghi"},
		// {
		// 	"double_digits",
		// 	args{"10#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"},
		// 	"jklmnopqrstuvwxyz",
		// },
		// {"mixed", args{"1 2 3 10#11#12#"}, "abcjkl"},
		// {"edge1", args{"1"}, "a"},
		// {"edge2", args{"26#"}, "z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := freqAlphabets(tt.args.s); got != tt.want {
				t.Errorf("freqAlphabets() = %v, want %v", got, tt.want)
			}
		})
	}
}
