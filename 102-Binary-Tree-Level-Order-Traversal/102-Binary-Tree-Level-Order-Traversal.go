package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int, 0)

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		subArray := []int{}
		for _, node := range queue {
			subArray = append(subArray, node.Val)
		}
		ans = append(ans, subArray)

		n := len(queue)
		for _, element := range queue {
			if element.Left != nil {
				queue = append(queue, element.Left)
			}
			if element.Right != nil {
				queue = append(queue, element.Right)
			}
		}
		queue = queue[n:]
	}

	return ans
}
