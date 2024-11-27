package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
  if p == nil || q == nil {
    return p == q
  }

	if p.Val != q.Val {
		return false
	}

	right := isSameTree(p.Right, q.Right)
	left := isSameTree(p.Left, q.Left)

	return right && left
}
