package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	runningSum := 0
	preorder(root, &runningSum, &sum)

	return sum
}

func preorder(root *TreeNode, runningSum *int, sum *int) {
	if root == nil {
		return
	}

	*runningSum = (*runningSum * 10) + root.Val

	if root.Left == nil && root.Right == nil {
		*sum += *runningSum
		*runningSum /= 10
		return
	}

	preorder(root.Left, runningSum, sum)
	preorder(root.Right, runningSum, sum)
	*runningSum /= 10
}
