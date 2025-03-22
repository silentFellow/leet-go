package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	ans := 0
	for _, child := range root.Children {
		ans = max(ans, maxDepth(child))
	}

	return 1 + ans
}
