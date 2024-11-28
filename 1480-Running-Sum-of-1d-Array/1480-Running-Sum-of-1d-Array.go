package leetcode

func runningSum(nums []int) []int {
	sum := 0

	for i, val := range nums {
		nums[i] = nums[i] + sum
		sum += val
	}

	return nums
}
