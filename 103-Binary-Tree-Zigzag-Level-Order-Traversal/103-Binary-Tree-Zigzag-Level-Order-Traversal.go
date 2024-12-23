package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
  if root == nil {
    return [][]int{}
  }

  ans := [][]int{}

  queue := []*TreeNode{root}
  level := 0

  for len(queue) > 0 {
    values := []int{}
    for _, node := range queue {
      values = append(values, node.Val)
    }

    if level % 2 != 0 {
      l, r := 0, len(queue)-1
      for l < r {
        values[l], values[r] = values[r], values[l]
        l++
        r--
      }
    }

    ans = append(ans, values)
    level++

    n := len(queue)
    for _, node := range queue {
      if node.Left != nil { queue = append(queue, node.Left) }
      if node.Right != nil { queue = append(queue, node.Right) }
    }
    queue = queue[n:]
  }


  return ans
}
