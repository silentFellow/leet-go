package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0

	var dfs func(cur *TreeNode) int
	dfs = func(cur *TreeNode) int {
		if cur == nil {
			return 0
		}

		left := dfs(cur.Left)
		right := dfs(cur.Right)

		ans = max(ans, left+right)

		return max(left, right) + 1
	}

	dfs(root)

	return ans
}
