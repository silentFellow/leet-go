package leetcode

import "testing"

func Test_secondHighest(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{s: "dfa12321afd"},
			want: 2,
		},
		{
			name: "example2",
			args: args{s: "abc1111"},
			want: -1,
		},
		{
			name: "no digits",
			args: args{s: "abcdef"},
			want: -1,
		},
		{
			name: "single digit",
			args: args{s: "a1b"},
			want: -1,
		},
		{
			name: "multiple digits",
			args: args{s: "a1b2c3d4"},
			want: 3,
		},
		{
			name: "same digits",
			args: args{s: "11122333"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondHighest(tt.args.s); got != tt.want {
				t.Errorf("secondHighest() = %v, want %v", got, tt.want)
			}
		})
	}
}
