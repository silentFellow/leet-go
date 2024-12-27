package leetcode

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{nums: []int{1, 2, 3, 4}},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "example2",
			args: args{nums: []int{-1, 1, 0, -3, 3}},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "all_ones",
			args: args{nums: []int{1, 1, 1, 1}},
			want: []int{1, 1, 1, 1},
		},
		{
			name: "single_zero",
			args: args{nums: []int{1, 2, 0, 4}},
			want: []int{0, 0, 8, 0},
		},
		{
			name: "negative_numbers",
			args: args{nums: []int{-1, -2, -3, -4}},
			want: []int{-24, -12, -8, -6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
