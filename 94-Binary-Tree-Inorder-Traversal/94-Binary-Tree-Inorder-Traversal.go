package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}
	current := root

	for current != nil || len(stack) != 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		current = current.Right
	}

	return result
}

// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
//
// 	nums := make([]int, 0)
//
// 	nums = append(nums, inorderTraversal(root.Left)...)
// 	nums = append(nums, root.Val)
// 	nums = append(nums, inorderTraversal(root.Right)...)
//
// 	return nums
// }
