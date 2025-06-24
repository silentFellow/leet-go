package leetcode

import (
	"reflect"
	"testing"
)

func Test_findKDistantIndices(t *testing.T) {
	type args struct {
		nums []int
		key  int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{nums: []int{3, 4, 9, 1, 3, 9, 5}, key: 9, k: 1},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "example2",
			args: args{nums: []int{2, 2, 2, 2, 2}, key: 2, k: 2},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "no k-distant indices",
			args: args{nums: []int{1, 3, 5, 7}, key: 2, k: 1},
			want: []int{},
		},
		{
			name: "key at start and end",
			args: args{nums: []int{5, 1, 1, 1, 5}, key: 5, k: 2},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "single element array",
			args: args{nums: []int{7}, key: 7, k: 1},
			want: []int{0},
		},
		{
			name: "k is zero",
			args: args{nums: []int{1, 2, 3, 2, 1}, key: 2, k: 0},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKDistantIndices(tt.args.nums, tt.args.key, tt.args.k); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("findKDistantIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
