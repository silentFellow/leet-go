package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeAnagrams(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{words: []string{"abba", "baba", "bbaa", "cd", "cd"}},
			want: []string{"abba", "cd"},
		},
		{
			name: "Example 2",
			args: args{words: []string{"a", "b", "c", "d", "e"}},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "No anagrams, repeated words",
			args: args{words: []string{"a", "a", "a"}},
			want: []string{"a"},
		},
		{
			name: "Non-consecutive anagrams",
			args: args{words: []string{"a", "b", "a"}},
			want: []string{"a", "b", "a"},
		},
		{
			name: "Multiple consecutive anagrams",
			args: args{words: []string{"z", "z", "z", "gsw", "wsg", "gsw", "krptu"}},
			want: []string{"z", "gsw", "krptu"},
		},
		{
			name: "Mixed anagrams and non-anagrams",
			args: args{words: []string{"ab", "bc", "cb", "ab", "bc"}},
			want: []string{"ab", "bc", "ab", "bc"},
		},
		{
			name: "Single word",
			args: args{words: []string{"abc"}},
			want: []string{"abc"},
		},
		{
			name: "Empty input",
			args: args{words: []string{}},
			want: []string{},
		},
		{
			name: "Two anagrams",
			args: args{words: []string{"listen", "silent"}},
			want: []string{"listen"},
		},
		{
			name: "Longer words with anagrams",
			args: args{words: []string{"abbb", "aaab"}},
			want: []string{"abbb", "aaab"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeAnagrams(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
