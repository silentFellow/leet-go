package leetcode

import "testing"

func Test_countPrefixes(t *testing.T) {
	type args struct {
		words []string
		s     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_case_1",
			args: args{
				words: []string{"a", "b", "c", "ab", "bc", "abc"},
				s:     "abc",
			},
			want: 3,
		},
		{
			name: "test_case_2",
			args: args{
				words: []string{"a", "a"},
				s:     "aa",
			},
			want: 2,
		},
		{
			name: "test_case_4",
			args: args{
				words: []string{"prefix", "pre", "fix"},
				s:     "prefixation",
			},
			want: 2,
		},
		{
			name: "test_case_5",
			args: args{
				words: []string{"a", "aa", "aaa"},
				s:     "aaaa",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrefixes(tt.args.words, tt.args.s); got != tt.want {
				t.Errorf("countPrefixes() = %v, want %v", got, tt.want)
			}
		})
	}
}
