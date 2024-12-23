package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
  if root == nil {
    return [][]int{}
  }

  ans := [][]int{}

  queue := []*TreeNode{root}
  for len(queue) > 0 {
    values := []int{}
    for _, node := range queue {
      values = append(values, node.Val)
    }
    ans = append(ans, values)

    n := len(queue)
    for _, node := range queue {
      if node.Left != nil { queue = append(queue, node.Left) }
      if node.Right != nil { queue = append(queue, node.Right) }
    }
    queue = queue[n:]
  }

  // reverse
  l, r := 0, len(ans)-1
  for l < r {
    ans[l], ans[r] = ans[r], ans[l]
    l++
    r--
  }

  return ans
}
