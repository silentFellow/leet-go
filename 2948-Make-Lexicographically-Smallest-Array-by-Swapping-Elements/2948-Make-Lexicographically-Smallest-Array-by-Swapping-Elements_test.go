package leetcode

import (
	"reflect"
	"testing"
)

func Test_lexicographicallySmallestArray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{nums: []int{1, 5, 3, 9, 8}, limit: 2},
			want: []int{1, 3, 5, 8, 9},
		},
		{
			name: "example2",
			args: args{nums: []int{1, 7, 6, 18, 2, 1}, limit: 3},
			want: []int{1, 6, 7, 18, 1, 2},
		},
		{
			name: "example3",
			args: args{nums: []int{1, 7, 28, 19, 10}, limit: 3},
			want: []int{1, 7, 28, 19, 10},
		},
		{
			name: "single_element",
			args: args{nums: []int{5}, limit: 1},
			want: []int{5},
		},
		{
			name: "no_swaps_needed",
			args: args{nums: []int{1, 2, 3, 4, 5}, limit: 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "all_elements_same",
			args: args{nums: []int{2, 2, 2, 2}, limit: 1},
			want: []int{2, 2, 2, 2},
		},
		{
			name: "large_limit",
			args: args{nums: []int{4, 3, 2, 1}, limit: 10},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lexicographicallySmallestArray(tt.args.nums, tt.args.limit); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("lexicographicallySmallestArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
