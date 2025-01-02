package leetcode

import (
	"reflect"
	"testing"
)

func Test_vowelStrings(t *testing.T) {
	type args struct {
		words   []string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test_case_1",
			args: args{
				words:   []string{"aba", "bcb", "ece", "aa", "e"},
				queries: [][]int{{0, 2}, {1, 4}, {1, 1}},
			},
			want: []int{2, 3, 0},
		},
		{
			name: "test_case_2",
			args: args{
				words:   []string{"a", "e", "i"},
				queries: [][]int{{0, 2}, {0, 1}, {2, 2}},
			},
			want: []int{3, 2, 1},
		},
		{
			name: "test_case_4",
			args: args{
				words:   []string{"xyz", "abc", "def", "ghi"},
				queries: [][]int{{0, 3}, {1, 2}, {2, 3}},
			},
			want: []int{0, 0, 0},
		},
		{
			name: "test_case_5",
			args: args{
				words:   []string{"aeiou", "iouea", "uoiea"},
				queries: [][]int{{0, 2}, {1, 1}, {2, 2}},
			},
			want: []int{3, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vowelStrings(tt.args.words, tt.args.queries); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("vowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
