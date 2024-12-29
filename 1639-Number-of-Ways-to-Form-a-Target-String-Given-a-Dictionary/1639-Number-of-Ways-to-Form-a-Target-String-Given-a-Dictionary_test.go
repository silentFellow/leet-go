package leetcode

import "testing"

func Test_numWays(t *testing.T) {
	type args struct {
		words  []string
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				words:  []string{"acca", "bbbb", "caca"},
				target: "aba",
			},
			want: 6,
		},
		{
			name: "example2",
			args: args{
				words:  []string{"abba", "baab"},
				target: "bab",
			},
			want: 4,
		},
		{
			name: "single_word",
			args: args{
				words:  []string{"abc"},
				target: "abc",
			},
			want: 1,
		},
		{
			name: "no_possible_way",
			args: args{
				words:  []string{"abc", "def"},
				target: "ghi",
			},
			want: 0,
		},
		{
			name: "multiple_ways",
			args: args{
				words:  []string{"aaa", "aaa", "aaa"},
				target: "aaa",
			},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.words, tt.args.target); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
