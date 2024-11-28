package leetcode

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}},
			want: []int{3, 2, 1},
		},
		{
			name: "Example 2",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 7},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 9},
					},
				},
			}},
			want: []int{4, 6, 7, 5, 2, 9, 8, 3, 1},
		},
		{
			name: "Example 3",
			args: args{root: nil},
			want: []int{},
		},
		{
			name: "Example 4",
			args: args{root: &TreeNode{Val: 1}},
			want: []int{1},
		},
		{
			name: "Single node with left child",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}},
			want: []int{2, 1},
		},
		{
			name: "Single node with right child",
			args: args{root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
