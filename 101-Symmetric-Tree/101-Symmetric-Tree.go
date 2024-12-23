package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
  queue := []*TreeNode{root}

  for len(queue) > 0 {
    values := []int{}
    for _, node := range queue {
      if node == nil {
        values = append(values, 999) // temp val, constraints -100 to 100
        continue
      }
      values = append(values, node.Val)
    }
    if !checkSymetric(values) {
      return false
    }

    n := len(queue)
    for _, node := range queue {
      if node != nil {
      queue = append(queue, node.Left)
      queue = append(queue, node.Right)
      }
    }
    queue = queue[n:]
  }

  return true
}

func checkSymetric(nums []int) bool {
  l, r := 0, len(nums)-1
  for l < r {
    if nums[l] != nums[r] { return false }
    l++
    r--
  }

  return true
}
