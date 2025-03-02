package leetcode

import (
	"reflect"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test_case_1",
			args: args{
				s1: "this apple is sweet",
				s2: "this apple is sour",
			},
			want: []string{"sweet", "sour"},
		},
		{
			name: "test_case_2",
			args: args{
				s1: "apple apple",
				s2: "banana",
			},
			want: []string{"banana"},
		},
		{
			name: "test_case_4",
			args: args{
				s1: "hello world",
				s2: "hello",
			},
			want: []string{"world"},
		},
		{
			name: "test_case_5",
			args: args{
				s1: "a b c",
				s2: "a b c d",
			},
			want: []string{"d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uncommonFromSentences(tt.args.s1, tt.args.s2); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("uncommonFromSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}
