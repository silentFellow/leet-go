package leetcode

func xorAllNums(nums1 []int, nums2 []int) int {
	totalXorNum2 := 0
	for _, num := range nums2 {
		totalXorNum2 ^= num
	}

	isEven := len(nums2)%2 == 0

	ans := 0
	for _, num := range nums1 {
		val := totalXorNum2

		// is the len of nums2 is even
		// then elements from num1 cancels each other
		// [1], [2, 3] => (1 ^ 2 ^ 1 ^ 3) => (2 ^ 3) => totalXorNum2
		if !isEven {
			val ^= num
		}

		ans ^= val
	}

	return ans
}
