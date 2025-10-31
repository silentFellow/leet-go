package leetcode

import (
	"reflect"
	"testing"
)

func Test_getSneakyNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{0, 1, 1, 0}},
			want: []int{0, 1},
		},
		{
			name: "Example 2",
			args: args{nums: []int{0, 3, 2, 1, 3, 2}},
			want: []int{2, 3},
		},
		{
			name: "Example 3",
			args: args{nums: []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}},
			want: []int{4, 5},
		},
		{
			name: "Edge case n=2",
			args: args{nums: []int{0, 1, 0, 1}},
			want: []int{0, 1},
		},
		{
			name: "Edge case n=3",
			args: args{nums: []int{2, 0, 1, 2, 1}},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSneakyNumbers(tt.args.nums)
			gotSet := map[int]struct{}{got[0]: {}, got[1]: {}}
			wantSet := map[int]struct{}{tt.want[0]: {}, tt.want[1]: {}}
			if !reflect.DeepEqual(gotSet, wantSet) {
				t.Errorf("getSneakyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
