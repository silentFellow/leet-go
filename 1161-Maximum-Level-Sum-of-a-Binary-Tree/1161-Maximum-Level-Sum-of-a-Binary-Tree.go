package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type maxSum struct {
	val   int
	level int
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := maxSum{
		val:   root.Val,
		level: 1,
	}

	queue := []*TreeNode{root}
	level := 1

	for len(queue) > 0 {
		sum := 0
		for _, node := range queue {
			sum += node.Val
		}
		if sum > ans.val {
			ans.val = sum
			ans.level = level
		}

		n := len(queue)
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
		level++
	}

	return ans.level
}
