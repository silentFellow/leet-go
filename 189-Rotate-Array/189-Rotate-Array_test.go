package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
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
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			want: []int{3, 99, -1, -100},
		},
		{
			name: "No rotation",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    0,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Full rotation",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Rotation greater than length",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    7,
			},
			want: []int{4, 5, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
