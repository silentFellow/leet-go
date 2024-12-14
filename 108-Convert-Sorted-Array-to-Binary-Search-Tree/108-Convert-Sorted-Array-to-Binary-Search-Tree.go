package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	root := &TreeNode{Val: nums[mid]}

	root.Left = buildBST(nums, start, mid-1)
	root.Right = buildBST(nums, mid+1, end)

	return root
}
