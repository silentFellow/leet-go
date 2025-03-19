package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortByBits(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
			want: []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			name: "Example 2",
			args: args{arr: []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}},
			want: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
		{
			name: "All zeros",
			args: args{arr: []int{0, 0, 0}},
			want: []int{0, 0, 0},
		},
		{
			name: "Single element",
			args: args{arr: []int{7}},
			want: []int{7},
		},
		{
			name: "Mixed bits",
			args: args{arr: []int{3, 7, 8, 6, 5}},
			want: []int{8, 3, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByBits(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
