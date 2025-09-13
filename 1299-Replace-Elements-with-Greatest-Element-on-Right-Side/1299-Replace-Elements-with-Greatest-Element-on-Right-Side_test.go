package leetcode

import (
	"reflect"
	"testing"
)

func Test_replaceElements(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{arr: []int{17, 18, 5, 4, 6, 1}},
			want: []int{18, 6, 6, 6, 1, -1},
		},
		// {
		// 	name: "single element",
		// 	args: args{arr: []int{400}},
		// 	want: []int{-1},
		// },
		// {
		// 	name: "increasing order",
		// 	args: args{arr: []int{1, 2, 3, 4, 5}},
		// 	want: []int{5, 5, 5, 5, -1},
		// },
		// {
		// 	name: "decreasing order",
		// 	args: args{arr: []int{5, 4, 3, 2, 1}},
		// 	want: []int{4, 3, 2, 1, -1},
		// },
		// {
		// 	name: "all same",
		// 	args: args{arr: []int{7, 7, 7, 7}},
		// 	want: []int{7, 7, 7, -1},
		// },
		// {
		// 	name: "two elements",
		// 	args: args{arr: []int{9, 8}},
		// 	want: []int{8, -1},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceElements(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
