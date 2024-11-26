package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxI, maxV := findMaxWithIndex(nums)
	root := &TreeNode{
		Val: maxV,
	}

	root.Left = constructMaximumBinaryTree(nums[:maxI])
	root.Right = constructMaximumBinaryTree(nums[maxI+1:])

	return root
}

func findMaxWithIndex(nums []int) (int, int) {
	maxI, maxV := 0, nums[0]

	for i, val := range nums {
		if val > maxV {
			maxI, maxV = i, val
		}
	}

	return maxI, maxV
}
