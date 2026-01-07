package leetcode

const MOD = 1e9 + 7

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	totalSum := 0

	var sum func(root *TreeNode) int
	sum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		return root.Val + sum(root.Left) + sum(root.Right)
	}
	totalSum = sum(root)

	best := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		s := root.Val + dfs(root.Left) + dfs(root.Right)
		curSplit := s * (totalSum - s)
		best = max(best, curSplit)

		return s
	}
	dfs(root)

	return best % MOD
}
