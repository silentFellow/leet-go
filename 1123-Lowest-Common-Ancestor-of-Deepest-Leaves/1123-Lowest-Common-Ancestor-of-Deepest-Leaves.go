package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode, depth int) (*TreeNode, int)
	dfs = func(node *TreeNode, depth int) (*TreeNode, int) {
		if node == nil {
			return nil, depth
		}

		left_lca, left_depth := dfs(node.Left, depth+1)
		right_lca, right_depth := dfs(node.Right, depth+1)

		if left_depth > right_depth {
			return left_lca, left_depth
		} else if right_depth > left_depth {
			return right_lca, right_depth
		} else {
			return node, left_depth
		}
	}

	lca, _ := dfs(root, 0)
	return lca
}
