package leetcode

func findSmallestInteger(nums []int, value int) int {
	mexMap := make(map[int]int)

	for _, num := range nums {
		mod := ((num % value) + value) % value
		mexMap[mod]++
	}

	mex := 0
	for {
		r := mex % value
		count := mexMap[r]
		if count > 0 {
			mexMap[r]--
			mex++
		} else if count == 0 {
			return mex
		}
	}
}
