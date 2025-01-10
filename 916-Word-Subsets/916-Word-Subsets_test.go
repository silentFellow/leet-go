package leetcode

import (
	"reflect"
	"testing"
)

func Test_wordSubsets(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test_case_1",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"e", "o"},
			},
			want: []string{"facebook", "google", "leetcode"},
		},
		{
			name: "test_case_2",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"l", "e"},
			},
			want: []string{"apple", "google", "leetcode"},
		},
		{
			name: "test_case_4",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"z"},
			},
			want: []string{"amazon"},
		},
		{
			name: "test_case_5",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"x"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordSubsets(tt.args.words1, tt.args.words2); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("wordSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wordCounter(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{
			name: "test_case_1",
			args: args{word: "apple"},
			want: map[rune]int{'a': 1, 'p': 2, 'l': 1, 'e': 1},
		},
		{
			name: "test_case_2",
			args: args{word: "banana"},
			want: map[rune]int{'b': 1, 'a': 3, 'n': 2},
		},
		{
			name: "test_case_3",
			args: args{word: "google"},
			want: map[rune]int{'g': 2, 'o': 2, 'l': 1, 'e': 1},
		},
		{
			name: "test_case_4",
			args: args{word: "leetcode"},
			want: map[rune]int{'l': 1, 'e': 3, 't': 1, 'c': 1, 'o': 1, 'd': 1},
		},
		{
			name: "test_case_5",
			args: args{word: "amazon"},
			want: map[rune]int{'a': 2, 'm': 1, 'z': 1, 'o': 1, 'n': 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordCounter(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}
