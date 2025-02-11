package leetcode

import "testing"

func Test_removeOccurrences(t *testing.T) {
	type args struct {
		s    string
		part string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{s: "daabcbaabcbc", part: "abc"},
			want: "dab",
		},
		{
			name: "test2",
			args: args{s: "axxxxyyyyb", part: "xy"},
			want: "ab",
		},
		{
			name: "test3",
			args: args{s: "hello", part: "ll"},
			want: "heo",
		},
		{
			name: "test4",
			args: args{s: "aaaaa", part: "aa"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOccurrences(tt.args.s, tt.args.part); got != tt.want {
				t.Errorf("removeOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
