package leetcode

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example1",
			args: args{
				words: []string{"i", "love", "leetcode", "i", "love", "coding"},
				k:     2,
			},
			want: []string{"i", "love"},
		},
		{
			name: "example2",
			args: args{
				words: []string{
					"the",
					"day",
					"is",
					"sunny",
					"the",
					"the",
					"the",
					"sunny",
					"is",
					"is",
				},
				k: 4,
			},
			want: []string{"the", "is", "sunny", "day"},
		},
		{
			name: "single_word",
			args: args{
				words: []string{"hello"},
				k:     1,
			},
			want: []string{"hello"},
		},
		{
			name: "all_unique",
			args: args{
				words: []string{"a", "b", "c", "d"},
				k:     2,
			},
			want: []string{"a", "b"},
		},
		{
			name: "same_frequency",
			args: args{
				words: []string{"a", "b", "a", "b", "c", "c"},
				k:     2,
			},
			want: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.words, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
