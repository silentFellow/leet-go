package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1

	return int(math.Max(float64(left), float64(right)))
}
