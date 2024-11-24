package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{s: "42"}, want: 42},
		{name: "Example 2", args: args{s: "   -042"}, want: -42},
		{name: "Example 3", args: args{s: "1337c0d3"}, want: 1337},
		{name: "Example 4", args: args{s: "0-1"}, want: 0},
		{name: "Example 5", args: args{s: "words and 987"}, want: 0},
		{name: "Empty string", args: args{s: ""}, want: 0},
		{name: "Only whitespace", args: args{s: "    "}, want: 0},
		{name: "Positive number with sign", args: args{s: "+123"}, want: 123},
		{name: "Negative number with sign", args: args{s: "-123"}, want: -123},
		{name: "Number with leading zeros", args: args{s: "000123"}, want: 123},
		{name: "test", args: args{s: "20000000000000000000"}, want: 2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
