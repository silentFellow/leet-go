package leetcode

import (
	"reflect"
	"testing"
)

func Test_applyOperations(t *testing.T) {
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
			args: args{nums: []int{1, 2, 2, 1, 1, 0}},
			want: []int{1, 4, 2, 0, 0, 0},
		},
		{
			name: "example2",
			args: args{nums: []int{0, 1}},
			want: []int{1, 0},
		},
		{
			name: "no_operations_needed",
			args: args{nums: []int{1, 3, 5, 7}},
			want: []int{1, 3, 5, 7},
		},
		{
			name: "all_zeros",
			args: args{nums: []int{0, 0, 0, 0}},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "all_elements_same",
			args: args{nums: []int{2, 2, 2, 2}},
			want: []int{4, 4, 0, 0},
		},
		{
			name: "alternating_elements",
			args: args{nums: []int{1, 1, 2, 2, 3, 3}},
			want: []int{2, 4, 6, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyOperations(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
