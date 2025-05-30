package leetcode

// INTITUTION: use 2 kandine's algorithm
// one for min-sum subarray and another for max-sum subarray
// return max(abs(minSum), abs(maxSum))

func maxAbsoluteSum(nums []int) int {
	maxKandine, minKandine, maxSum, minSum := 0, 0, 0, 0

	for _, v := range nums {
		maxSum += v
		minSum += v

		if maxSum < 0 {
			maxSum = 0
		}
		if minSum > 0 {
			minSum = 0
		}

		maxKandine = max(maxKandine, maxSum)
		minKandine = min(minKandine, minSum)
	}

	return max(abs(maxKandine), abs(minKandine))
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
