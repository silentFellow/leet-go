package leetcode

func productExceptSelf(nums []int) []int {
	n := len(nums)

	prefixMul := make([]int, n)
	mul := 1
	for i, val := range nums {
		prefixMul[i] = mul
		mul *= val
	}

	postfixMul := make([]int, n)
	mul = 1
	for i := n - 1; i >= 0; i-- {
		postfixMul[i] = mul
		mul *= nums[i]
	}

	for i := range nums {
		nums[i] = prefixMul[i] * postfixMul[i]
	}

	return nums
}
