package leetcode

import "math"

func maxAscendingSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sum, maxSum := nums[0], 0

	for i, j := 0, 1; j < len(nums); i, j = i+1, j+1 {
		if nums[i] < nums[j] {
			sum += nums[j]
		} else {
			maxSum = int(math.Max(float64(sum), float64(maxSum)))
			sum = nums[j]
		}
	}

	maxSum = int(math.Max(float64(sum), float64(maxSum)))

	return maxSum
}
