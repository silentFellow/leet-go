package leetcode

func removeDuplicates(nums []int) int {
	k, prev := 1, nums[0]

	r := 1
	for r < len(nums) {
		for r < len(nums) && nums[r] == prev {
			r++
		}

		if r < len(nums) {
			prev = nums[r]
			nums[k] = nums[r]
			k++
		}
	}

	return k
}
