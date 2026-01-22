package leetcode

func minimumPairRemoval(nums []int) int {
	count := 0
	for {
		// Check if non-decreasing
		ok := true
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[i-1] {
				ok = false
				break
			}
		}
		if ok {
			return count
		}

		// Find leftmost minimum-sum adjacent pair
		minSum := nums[0] + nums[1]
		idx := 0
		for i := 1; i < len(nums)-1; i++ {
			sum := nums[i] + nums[i+1]
			if sum < minSum {
				minSum = sum
				idx = i
			}
		}

		// Merge the pair
		nums = append(nums[:idx], append([]int{nums[idx] + nums[idx+1]}, nums[idx+2:]...)...)
		count++
	}
}
