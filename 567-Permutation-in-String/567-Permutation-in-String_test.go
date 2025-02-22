package leetcode

import "testing"

func Test_checkInclusion(t *testing.T) {
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
			args: args{s1: "ab", s2: "eidbaooo"},
			want: true,
		},
		{
			name: "example2",
			args: args{s1: "ab", s2: "eidboaoo"},
			want: false,
		},
		{
			name: "empty_s1",
			args: args{s1: "", s2: "anything"},
			want: true,
		},
		{
			name: "empty_s2",
			args: args{s1: "a", s2: ""},
			want: false,
		},
		{
			name: "single_char_match",
			args: args{s1: "a", s2: "a"},
			want: true,
		},
		{
			name: "single_char_no_match",
			args: args{s1: "a", s2: "b"},
			want: false,
		},
		{
			name: "longer_s1",
			args: args{s1: "abcd", s2: "dcba"},
			want: true,
		},
		{
			name: "no_permutation",
			args: args{s1: "abc", s2: "defghijkl"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
