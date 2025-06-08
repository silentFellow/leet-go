package leetcode

import "testing"

func Test_clearStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"no stars", args{"abc"}, "abc"},
		{"one star, remove smallest left", args{"aaba*"}, "aab"},
		{"multiple stars", args{"cb*a*b*"}, "c"},
		{"all stars", args{"a*b*c*"}, ""},
		{"star at start", args{"*abc"}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clearStars(tt.args.s); got != tt.want {
				t.Errorf("clearStars() = %v, want %v", got, tt.want)
			}
		})
	}
}
