package leetcode

import (
	"reflect"
	"testing"
)

func Test_stringMatching(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{words: []string{"mass", "as", "hero", "superhero"}},
			want: []string{"as", "hero"},
		},
		{
			name: "test2",
			args: args{words: []string{"leetcode", "et", "code"}},
			want: []string{"et", "code"},
		},
		{
			name: "test3",
			args: args{words: []string{"blue", "green", "bu"}},
			want: []string{},
		},
		{
			name: "test4",
			args: args{words: []string{"apple", "app", "le", "ple"}},
			want: []string{"app", "le", "ple"},
		},
		{
			name: "test5",
			args: args{words: []string{"a", "b", "c", "abc"}},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringMatching(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
