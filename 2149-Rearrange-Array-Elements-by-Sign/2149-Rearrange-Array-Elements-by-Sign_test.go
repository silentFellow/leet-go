package leetcode

import (
	"reflect"
	"testing"
)

func Test_rearrangeArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{nums: []int{3, 1, -2, -5, 2, -4}},
			want: []int{3, -2, 1, -5, 2, -4},
		},
		{
			name: "test2",
			args: args{nums: []int{-1, 1}},
			want: []int{1, -1},
		},
		{
			name: "test3",
			args: args{nums: []int{1, -1, 2, -2, 3, -3}},
			want: []int{1, -1, 2, -2, 3, -3},
		},
		{
			name: "test4",
			args: args{nums: []int{-3, -1, -2, 5, 2, 4}},
			want: []int{5, -3, 2, -1, 4, -2},
		},
		{
			name: "test5",
			args: args{nums: []int{1, 2, 3, -1, -2, -3}},
			want: []int{1, -1, 2, -2, 3, -3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rearrangeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
