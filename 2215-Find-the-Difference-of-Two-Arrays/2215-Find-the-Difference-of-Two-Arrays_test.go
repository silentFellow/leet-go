package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func sortSlices(slices [][]int) {
	for _, slice := range slices {
		sort.Ints(slice)
	}
}

func normalizeSlices(slices [][]int) {
	for i, slice := range slices {
		if slice == nil {
			slices[i] = []int{}
		}
	}
}

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4, 6},
			},
			want: [][]int{{1, 3}, {4, 6}},
		},
		{
			name: "Example 2",
			args: args{
				nums1: []int{1, 2, 3, 3},
				nums2: []int{1, 1, 2, 2},
			},
			want: [][]int{{3}, {}},
		},
		{
			name: "No difference",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{1, 2, 3},
			},
			want: [][]int{{}, {}},
		},
		{
			name: "All different",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{4, 5, 6},
			},
			want: [][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			name: "Empty nums1",
			args: args{
				nums1: []int{},
				nums2: []int{1, 2, 3},
			},
			want: [][]int{{}, {1, 2, 3}},
		},
		{
			name: "Empty nums2",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{},
			},
			want: [][]int{{1, 2, 3}, {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifference(tt.args.nums1, tt.args.nums2)
			sortSlices(got)
			sortSlices(tt.want)
			normalizeSlices(got)
			normalizeSlices(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
