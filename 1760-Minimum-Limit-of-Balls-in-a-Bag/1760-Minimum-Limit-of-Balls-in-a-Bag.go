package leetcode

// TODO: verify once again

func minimumSize(nums []int, maxOperations int) int {
	sum := 0
	right := 0

	for _, num := range nums {
		sum += num
		right = max(num, right)
	}

	left := sum / (len(nums) + maxOperations)
	if sum % (len(nums)+maxOperations) != 0 {
		left++
	}

	for left < right {
		operations := 0
		mid := (left + right) / 2

		for _, num := range nums {
			operations += num / mid
			if num % mid == 0 {
				operations--
			}

			if operations > maxOperations {
				break
			}
		}

		if operations > maxOperations {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
