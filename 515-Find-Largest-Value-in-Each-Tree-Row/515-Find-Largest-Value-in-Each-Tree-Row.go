package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
  // edge case
	if root == nil {
		return []int{}
	}

  // create array to store ans
  // queue for BFS in a tree
	ans := []int{}
	queue := []*TreeNode{root}

  // loop till all nodes visited
	for len(queue) > 0 {
		maxv := queue[0].Val

    // find max in the level
		for _, node := range queue {
			maxv = max(maxv, node.Val)
		}
		ans = append(ans, maxv)

    // add next level
		n := len(queue)
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

    // remove the current level
		queue = queue[n:]
	}

  // return ans
	return ans
}
