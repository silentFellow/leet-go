package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	sum := 0

	ans := preOrder(root, &sum, targetSum)
	return ans
}

func preOrder(root *TreeNode, sum *int, target int) bool {
	if root == nil {
		return false
	}

	*sum += root.Val

	if root.Left == nil && root.Right == nil {
		isValid := *sum == target
		*sum -= root.Val
		return isValid
	}

	left := preOrder(root.Left, sum, target)
	right := preOrder(root.Right, sum, target)
	*sum -= root.Val

	return left || right
}
