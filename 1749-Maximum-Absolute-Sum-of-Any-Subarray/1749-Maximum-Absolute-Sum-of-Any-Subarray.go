package leetcode

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
