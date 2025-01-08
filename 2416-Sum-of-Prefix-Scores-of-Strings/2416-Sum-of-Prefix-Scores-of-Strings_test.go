package leetcode

import (
	"reflect"
	"testing"
)

func Test_sumPrefixScores(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{words: []string{"abc", "ab", "bc", "b"}},
			want: []int{5, 4, 3, 2},
		},
		{
			name: "example2",
			args: args{words: []string{"abcd"}},
			want: []int{4},
		},
		{
			name: "single character",
			args: args{words: []string{"a", "b", "c"}},
			want: []int{1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumPrefixScores(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumPrefixScores() = %v, want %v", got, tt.want)
			}
		})
	}
}
