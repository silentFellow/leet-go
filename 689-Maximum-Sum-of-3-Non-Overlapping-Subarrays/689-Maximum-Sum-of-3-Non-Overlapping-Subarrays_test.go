package leetcode

import (
	"reflect"
	"testing"
)

func Test_maxSumOfThreeSubarrays(t *testing.T) {
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
				nums: []int{1, 2, 1, 2, 6, 7, 5, 1},
				k:    2,
			},
			want: []int{0, 3, 5},
		},
		{
			name: "example2",
			args: args{
				nums: []int{1, 2, 1, 2, 1, 2, 1, 2, 1},
				k:    2,
			},
			want: []int{0, 2, 4},
		},
		{
			name: "single_element",
			args: args{
				nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1},
				k:    3,
			},
			want: []int{0, 3, 6},
		},
		{
			name: "all_same",
			args: args{
				nums: []int{5, 5, 5, 5, 5, 5, 5, 5, 5},
				k:    3,
			},
			want: []int{0, 3, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumOfThreeSubarrays(tt.args.nums, tt.args.k); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("maxSumOfThreeSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
