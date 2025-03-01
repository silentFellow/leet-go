package leetcode

func applyOperations(nums []int) []int {
	// apply the operation
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums[i-1] *= 2
			nums[i] = 0
		}
	}

	// move all non-zero elements to the end
	cur, next := 0, 0
	for next != len(nums) {
		for next != len(nums) && nums[next] == 0 {
			next++
		}

		if next != len(nums) {
			nums[cur] = nums[next]
			cur++
			next++
		}
	}

	for cur < len(nums) {
		nums[cur] = 0
		cur++
	}

	return nums
}
