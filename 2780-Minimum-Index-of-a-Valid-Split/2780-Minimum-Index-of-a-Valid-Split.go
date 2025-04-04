package leetcode

func minimumIndex(nums []int) int {
	// Step 1: Find the majority element using Boyer-Moore Voting Algorithm
	candidate, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	// Step 2: Count occurrences of the majority element
	totalCount := 0
	for _, num := range nums {
		if num == candidate {
			totalCount++
		}
	}

	// Step 3: Find the minimum index where the split is valid
	leftCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == candidate {
			leftCount++
		}
		rightCount := totalCount - leftCount

		if leftCount > (i+1)/2 && rightCount > (len(nums)-(i+1))/2 {
			return i
		}
	}

	return -1
}
