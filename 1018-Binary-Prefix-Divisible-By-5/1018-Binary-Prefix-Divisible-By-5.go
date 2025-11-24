package leetcode

func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))

	rem := 0
	for i, v := range nums {
		// rem*2 == (rem << 1)
		rem = (rem*2 + v) % 5
		ans[i] = rem == 0
	}

	return ans
}
