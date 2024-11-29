package leetcode

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
		{
			name: "example 2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: []int{-1, -1},
		},
		{
			name: "example 3",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: []int{-1, -1},
		},
		{
			name: "single element found",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: []int{0, 0},
		},
		{
			name: "single element not found",
			args: args{
				nums:   []int{1},
				target: 2,
			},
			want: []int{-1, -1},
		},
		{
			name: "multiple occurrences",
			args: args{
				nums:   []int{2, 2, 2, 2, 2},
				target: 2,
			},
			want: []int{0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
