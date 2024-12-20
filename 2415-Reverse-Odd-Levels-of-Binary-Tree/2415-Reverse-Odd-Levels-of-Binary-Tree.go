package leetcode

// INTITUTION: use BFS to traverse level wise
// if level is odd reverse ans reassign it

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		n := len(queue)
		levelValues := []int{}

		// collect all values of the level
		for i := range n {
			node := queue[i]
			levelValues = append(levelValues, node.Val)
		}

		// reverse if level is odd
		if level%2 == 1 {
			for i, j := 0, len(levelValues)-1; i < j; i, j = i+1, j-1 {
				levelValues[i], levelValues[j] = levelValues[j], levelValues[i]
			}
		}

		// reassign to the tree in reverse order
		for i := range n {
			node := queue[i]
			node.Val = levelValues[i]
		}

		// append the next level
		// with respect to queue
		for i := range n {
			node := queue[i]
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

	return root
}
