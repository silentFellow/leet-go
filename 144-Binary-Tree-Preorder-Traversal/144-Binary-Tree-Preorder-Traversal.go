package leetcode

// TODO: Follow up: Recursive solution is trivial, could you do it iteratively?

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := make([]int, 0)

	nums = append(nums, root.Val)
	nums = append(nums, preorderTraversal(root.Left)...)
	nums = append(nums, preorderTraversal(root.Right)...)

	return nums
}
