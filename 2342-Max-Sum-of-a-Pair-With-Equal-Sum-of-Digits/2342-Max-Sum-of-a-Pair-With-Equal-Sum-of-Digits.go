package leetcode

func maximumSum(nums []int) int {
	hmap := make(map[int]int)

	ans := -1
	for _, val := range nums {
		sum := sumOfDigit(val)
		if v, ok := hmap[sum]; ok {
			ans = max(ans, val+v)
		}

		hmap[sum] = max(hmap[sum], val)
	}

	return ans
}

// calculate sum of digits
func sumOfDigit(v int) int {
	ans := 0

	for v != 0 {
		ans += v % 10
		v /= 10
	}

	return ans
}
