package leetcode

import (
	"reflect"
	"testing"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "example 1",
			args: args{nums: []int{3, 2, 1, 6, 0, 5}},
			want: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 0,
					},
				},
			},
		},
		{
			name: "example 2",
			args: args{nums: []int{3, 2, 1}},
			want: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
		{
			name: "single element",
			args: args{nums: []int{1}},
			want: &TreeNode{
				Val: 1,
			},
		},
		{
			name: "two elements",
			args: args{nums: []int{2, 1}},
			want: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		{
			name: "empty array",
			args: args{nums: []int{}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
