package leetcode

import (
	"testing"
)

func isValidBST(root *TreeNode, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if (min != nil && root.Val <= min.Val) || (max != nil && root.Val >= max.Val) {
		return false
	}
	return isValidBST(root.Left, min, root) && isValidBST(root.Right, root, max)
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := height(root.Right)
	if rightHeight == -1 {
		return -1
	}
	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "example1",
			args: args{nums: []int{-10, -3, 0, 5, 9}},
		},
		{
			name: "example2",
			args: args{nums: []int{1, 3}},
		},
		{
			name: "single element",
			args: args{nums: []int{1}},
		},
		{
			name: "empty array",
			args: args{nums: []int{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedArrayToBST(tt.args.nums)
			if !isValidBST(got, nil, nil) {
				t.Errorf("sortedArrayToBST() produced an invalid BST")
			}
			if !isBalanced(got) {
				t.Errorf("sortedArrayToBST() produced an unbalanced BST")
			}
		})
	}
}
