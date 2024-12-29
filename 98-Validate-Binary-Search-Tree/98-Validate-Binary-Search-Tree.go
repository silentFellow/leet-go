package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
  return Validate(root, math.MinInt64, math.MaxInt64)
}

func Validate(node *TreeNode, min int, max int) bool {
  if node == nil {
    return true
  }

  if node.Val <= min || node.Val >= max {
    return false
  }

  return Validate(node.Left, min, node.Val) && Validate(node.Right, node.Val, max)
}
