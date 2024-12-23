package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
  if root == nil {
    return [][]int{}
  }

	ans := [][]int{}

	queue := []*Node{root}
	for len(queue) > 0 {
		values := []int{}
		for _, node := range queue {
			values = append(values, node.Val)
		}
		ans = append(ans, values)

		n := len(queue)
		for _, node := range queue {
			for _, child := range node.Children {
				if child != nil {
					queue = append(queue, child)
				}
			}
		}
		queue = queue[n:]
	}

	return ans
}
