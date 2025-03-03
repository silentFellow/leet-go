package leetcode

// TODO: try to do in constant space

func pivotArray(nums []int, pivot int) []int {
	ans := make([]int, len(nums))
	pivotCount := 0
	idx := 0

	for _, v := range nums {
		if v < pivot {
			ans[idx] = v
			idx++
		}

		if v == pivot {
			pivotCount++
		}
	}

	for range pivotCount {
		ans[idx] = pivot
		idx++
	}

	for _, v := range nums {
		if v > pivot {
			ans[idx] = v
			idx++
		}
	}

  return ans
}
