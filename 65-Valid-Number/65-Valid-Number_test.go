package leetcode

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example 1", args: args{s: "0"}, want: true},
		{name: "Example 2", args: args{s: "e"}, want: false},
		{name: "Example 3", args: args{s: "."}, want: false},
		{name: "Empty string", args: args{s: ""}, want: false},
		{name: "Valid integer", args: args{s: "123"}, want: true},
		{name: "Valid float", args: args{s: "3.14"}, want: true},
		{name: "Invalid number with letters", args: args{s: "1a"}, want: false},
		{name: "Valid number with exponent", args: args{s: "2e10"}, want: true},
		{name: "Invalid number with multiple dots", args: args{s: "1.2.3"}, want: false},
		{name: "Valid negative number", args: args{s: "-123"}, want: true},
		{name: "Valid out of range number", args: args{s: "-8115e957"}, want: true},
		{name: "Invalid out of range number", args: args{s: "-8115e957e"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
