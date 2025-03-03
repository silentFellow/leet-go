package leetcode

import (
	"reflect"
	"testing"
)

func Test_pivotArray(t *testing.T) {
	type args struct {
		nums  []int
		pivot int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				nums:  []int{9, 12, 5, 10, 14, 3, 10},
				pivot: 10,
			},
			want: []int{9, 5, 3, 10, 10, 12, 14},
		},
		{
			name: "example2",
			args: args{
				nums:  []int{-3, 4, 3, 2},
				pivot: 2,
			},
			want: []int{-3, 2, 4, 3},
		},
		{
			name: "all_less_than_pivot",
			args: args{
				nums:  []int{1, 2, 3},
				pivot: 5,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "all_greater_than_pivot",
			args: args{
				nums:  []int{6, 7, 8},
				pivot: 5,
			},
			want: []int{6, 7, 8},
		},
		{
			name: "mixed_elements",
			args: args{
				nums:  []int{5, 1, 7, 2, 6, 3},
				pivot: 4,
			},
			want: []int{1, 2, 3, 5, 7, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotArray(tt.args.nums, tt.args.pivot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pivotArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
