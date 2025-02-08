package leetcode

func maxAscendingSum(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	sum, maxSum := nums[0], 0
	for i, j := 0, 1; j < len(nums); i, j = i+1, j+1 {
		if nums[j] > nums[i] {
			sum += nums[j]
		} else {
			maxSum = max(maxSum, sum)
			sum = nums[j]
		}
	}
	maxSum = max(maxSum, sum)

	return maxSum
}
