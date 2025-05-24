package leetcode

import (
	"reflect"
	"testing"
)

func Test_findWordsContaining(t *testing.T) {
	type args struct {
		words []string
		x     byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				words: []string{"leet", "code"},
				x:     'e',
			},
			want: []int{0, 1},
		},
		{
			name: "example2",
			args: args{
				words: []string{"abc", "bcd", "aaaa", "cbc"},
				x:     'a',
			},
			want: []int{0, 2},
		},
		{
			name: "no match",
			args: args{
				words: []string{"foo", "bar", "baz"},
				x:     'x',
			},
			want: []int{},
		},
		{
			name: "all match",
			args: args{
				words: []string{"a", "ba", "ca"},
				x:     'a',
			},
			want: []int{0, 1, 2},
		},
		{
			name: "empty words",
			args: args{
				words: []string{},
				x:     'z',
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWordsContaining(tt.args.words, tt.args.x); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("findWordsContaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
