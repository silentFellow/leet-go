package leetcode

import (
	"reflect"
	"testing"
)

func Test_divideString(t *testing.T) {
	type args struct {
		s    string
		k    int
		fill byte
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "all groups filled",
			args: args{s: "abcdefghi", k: 3, fill: 'x'},
			want: []string{"abc", "def", "ghi"},
		},
		{
			name: "last group needs fill",
			args: args{s: "abcdefghij", k: 3, fill: 'x'},
			want: []string{"abc", "def", "ghi", "jxx"},
		},
		{
			name: "single group, needs fill",
			args: args{s: "a", k: 3, fill: 'z'},
			want: []string{"azz"},
		},
		{
			name: "single group, no fill",
			args: args{s: "abc", k: 3, fill: 'y'},
			want: []string{"abc"},
		},
		{
			name: "multiple groups, last group needs fill",
			args: args{s: "abcdefg", k: 2, fill: 'q'},
			want: []string{"ab", "cd", "ef", "gq"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideString(tt.args.s, tt.args.k, tt.args.fill); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("divideString() = %v, want %v", got, tt.want)
			}
		})
	}
}
