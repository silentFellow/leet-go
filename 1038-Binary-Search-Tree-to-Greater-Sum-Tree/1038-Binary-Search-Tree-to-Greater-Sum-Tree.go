package leetcode

// Note: This question is the same as 538: https://leetcode.com/problems/convert-bst-to-greater-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
  sum := 0
  getSum(root, &sum)
  return root
}

func getSum(root *TreeNode, sum *int) {
  if root == nil {
    return
  }

  getSum(root.Right, sum)

  root.Val += *sum
  *sum = root.Val

  getSum(root.Left, sum)
}
