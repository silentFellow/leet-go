package leetcode

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "example2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "example3",
			args: args{
				nums: []int{9, 11},
				k:    2,
			},
			want: []int{11},
		},
		{
			name: "example4",
			args: args{
				nums: []int{4, -2},
				k:    2,
			},
			want: []int{4},
		},
		{
			name: "example5",
			args: args{
				nums: []int{1, 3, 1, 2, 0, 5},
				k:    3,
			},
			want: []int{3, 3, 2, 5},
		},
		{
			name: "example5",
			args: args{
				nums: []int{1, -1},
				k:    1,
			},
			want: []int{1, -1},
		},
		{
			name: "example5",
			args: args{
				nums: []int{9, 10, 9, -7, -4, -8, 2, -6},
				k:    5,
			},
			want: []int{10, 10, 9, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
