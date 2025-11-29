package leetcode

func minOperations(nums []int, k int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	ans := 0
	for sum % k != 0 {
		sum--
		ans++
	}

	return ans
}
