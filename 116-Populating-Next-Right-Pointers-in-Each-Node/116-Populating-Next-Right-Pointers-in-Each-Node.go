package leetcode

// TODO: try to reduce time complexity, current h*m
// Follow-up:
// You may only use constant extra space.
// The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		for i := len(queue) - 2; i >= 0; i-- {
			queue[i].Next = queue[i+1]
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
	}

	return root
}
