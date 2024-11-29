package leetcode

func numSubarrayProductLessThanK(nums []int, k int) int {
	product := 1
	count, l := 0, 0

	for r := range nums {
		product *= nums[r]

		for product >= k && l <= r {
			product /= nums[l]
			l++
		}
		count += (r - l) + 1
	}

	return count
}
