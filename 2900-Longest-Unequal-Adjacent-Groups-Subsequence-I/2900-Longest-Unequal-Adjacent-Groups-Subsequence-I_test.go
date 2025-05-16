package leetcode

import (
	"reflect"
	"testing"
)

func Test_getLongestSubsequence(t *testing.T) {
	type args struct {
		words  []string
		groups []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example1",
			args: args{
				words:  []string{"e", "a", "b"},
				groups: []int{0, 0, 1},
			},
			want: []string{"e", "b"},
		},
		{
			name: "example2",
			args: args{
				words:  []string{"a", "b", "c", "d"},
				groups: []int{1, 0, 1, 1},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "all_same_group",
			args: args{
				words:  []string{"x", "y", "z"},
				groups: []int{1, 1, 1},
			},
			want: []string{"x"},
		},
		{
			name: "alternating_groups",
			args: args{
				words:  []string{"a", "b", "c", "d", "e"},
				groups: []int{0, 1, 0, 1, 0},
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "two_elements",
			args: args{
				words:  []string{"foo", "bar"},
				groups: []int{1, 0},
			},
			want: []string{"foo", "bar"},
		},
		{
			name: "single_element",
			args: args{
				words:  []string{"solo"},
				groups: []int{0},
			},
			want: []string{"solo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLongestSubsequence(tt.args.words, tt.args.groups); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("getLongestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
