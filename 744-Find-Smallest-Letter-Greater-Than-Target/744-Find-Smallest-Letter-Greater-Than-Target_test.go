package leetcode

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Example 1",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'a'},
			want: 'c',
		},
		{
			name: "Example 2",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'c'},
			want: 'f',
		},
		{
			name: "Example 3",
			args: args{letters: []byte{'x', 'x', 'y', 'y'}, target: 'z'},
			want: 'x',
		},
		{
			name: "Multiple same letters",
			args: args{
				letters: []byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'},
				target:  'e',
			},
			want: 'n',
		},
		{
			name: "Target not present, between letters",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'd'},
			want: 'f',
		},
		{
			name: "Target before all letters",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'a'},
			want: 'c',
		},
		{
			name: "Target after all letters",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'k'},
			want: 'c',
		},
		{
			name: "Consecutive duplicates",
			args: args{letters: []byte{'a', 'b', 'b', 'b', 'c', 'c', 'd'}, target: 'b'},
			want: 'c',
		},
		{
			name: "Target is last letter",
			args: args{letters: []byte{'c', 'f', 'j', 'k', 'm', 'p'}, target: 'p'},
			want: 'c',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
