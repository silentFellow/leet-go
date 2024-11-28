package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := []int{}

	nums = append(nums, postorderTraversal(root.Left)...)
	nums = append(nums, postorderTraversal(root.Right)...)
	nums = append(nums, root.Val)

	return nums
}
