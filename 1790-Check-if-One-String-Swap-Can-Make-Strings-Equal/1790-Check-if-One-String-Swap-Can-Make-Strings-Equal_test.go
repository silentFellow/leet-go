package leetcode

import "testing"

func Test_areAlmostEqual(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{s1: "bank", s2: "kanb"},
			want: true,
		},
		{
			name: "example2",
			args: args{s1: "attack", s2: "defend"},
			want: false,
		},
		{
			name: "example3",
			args: args{s1: "kelb", s2: "kelb"},
			want: true,
		},
		{
			name: "example4",
			args: args{s1: "abcd", s2: "abdc"},
			want: true,
		},
		{
			name: "example5",
			args: args{s1: "abcd", s2: "abcd"},
			want: true,
		},
		{
			name: "example6",
			args: args{s1: "abcd", s2: "abcf"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areAlmostEqual(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("areAlmostEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
