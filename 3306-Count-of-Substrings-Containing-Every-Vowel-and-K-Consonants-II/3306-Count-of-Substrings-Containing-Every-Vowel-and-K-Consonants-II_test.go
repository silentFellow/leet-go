package leetcode

import "testing"

func Test_countOfSubstrings(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test_case_1",
			args: args{word: "aeioqq", k: 1},
			want: 0,
		},
		{
			name: "test_case_2",
			args: args{word: "aeiou", k: 0},
			want: 1,
		},
		{
			name: "test_case_3",
			args: args{word: "ieaouqqieaouqq", k: 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfSubstrings(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("countOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isvowel(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isvowel(tt.args.c); got != tt.want {
				t.Errorf("isvowel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countK(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countK(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("countK() = %v, want %v", got, tt.want)
			}
		})
	}
}
