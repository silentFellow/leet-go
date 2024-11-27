package leetcode

import "testing"

func Test_toLowerCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test 1", args: args{s: "Hello"}, want: "hello"},
		{name: "Test 2", args: args{s: "here"}, want: "here"},
		{name: "Test 3", args: args{s: "LOVELY"}, want: "lovely"},
		{name: "Test 4", args: args{s: "GoLang"}, want: "golang"},
		{name: "Test 5", args: args{s: "123ABC"}, want: "123abc"},
		{name: "Test 6", args: args{s: "MixedCASE"}, want: "mixedcase"},
		{name: "Test 7", args: args{s: ""}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.args.s); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
