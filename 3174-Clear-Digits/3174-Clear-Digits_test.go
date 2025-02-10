package leetcode

import "testing"

func Test_clearDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no digits",
			args: args{s: "abc"},
			want: "abc",
		},
		{
			name: "all digits",
			args: args{s: "1234"},
			want: "",
		},
		{
			name: "mixed characters",
			args: args{s: "a1b2c3"},
			want: "",
		},
		{
			name: "digits at the end",
			args: args{s: "abc123"},
			want: "",
		},
		{
			name: "digits at the start",
			args: args{s: "1a2b3c"},
			want: "c",
		},
		{
			name: "single digit",
			args: args{s: "a1"},
			want: "",
		},
		{
			name: "consecutive digits",
			args: args{s: "ab12cd34"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clearDigits(tt.args.s); got != tt.want {
				t.Errorf("clearDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
