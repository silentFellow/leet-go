package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
  if root == nil {
    return []float64{}
  }

  ans := []float64{}

  queue := []*TreeNode{root}
  for len(queue) > 0 {
    sum := 0
    n := len(queue)

    for _, node := range queue {
      sum += node.Val
    }
    avg := float64(sum) / float64(n)
    ans = append(ans, avg)

    for _, node := range queue {
      if node.Left != nil { queue = append(queue, node.Left) }
      if node.Right != nil { queue = append(queue, node.Right) }
    }
    queue = queue[n:]
  }

  return ans
}
