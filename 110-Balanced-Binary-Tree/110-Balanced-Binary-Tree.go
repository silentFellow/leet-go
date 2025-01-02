package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := getDepth(root.Left)
	right := getDepth(root.Right)

	if abs(left-right) <= 1 {
		left := isBalanced(root.Left)
		right := isBalanced(root.Right)
		return left && right
	} else {
		return false
	}
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getDepth(root.Left)
	right := getDepth(root.Right)

	return max(left, right) + 1
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
