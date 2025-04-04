package leetcode

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{s: "ababcbacadefegdehijhklij"},
			want: []int{9, 7, 8},
		},
		{
			name: "example2",
			args: args{s: "eccbbbbdec"},
			want: []int{10},
		},
		{
			name: "single_char",
			args: args{s: "a"},
			want: []int{1},
		},
		{
			name: "all_unique",
			args: args{s: "abcdef"},
			want: []int{1, 1, 1, 1, 1, 1},
		},
		{
			name: "repeated_chars",
			args: args{s: "aaabbbccc"},
			want: []int{3, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLabels(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
