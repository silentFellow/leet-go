package leetcode

import "testing"

func Test_possibleStringCount(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{word: "aabbccdd", k: 7},
			want: 5,
		},
		// {
		// 	name: "Example 2",
		// 	args: args{word: "aabbccdd", k: 8},
		// 	want: 1,
		// },
		// {
		// 	name: "Example 3",
		// 	args: args{word: "aaabbb", k: 3},
		// 	want: 8,
		// },
		// {
		// 	name: "Single char, k=1",
		// 	args: args{word: "a", k: 1},
		// 	want: 1,
		// },
		// {
		// 	name: "All unique, k=word.len",
		// 	args: args{word: "abcdef", k: 6},
		// 	want: 1,
		// },
		// {
		// 	name: "All unique, k=3",
		// 	args: args{word: "abcdef", k: 3},
		// 	want: 20,
		// },
		// {
		// 	name: "All same, k=1",
		// 	args: args{word: "aaaaa", k: 1},
		// 	want: 5,
		// },
		// {
		// 	name: "All same, k=5",
		// 	args: args{word: "aaaaa", k: 5},
		// 	want: 1,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleStringCount(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("possibleStringCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
