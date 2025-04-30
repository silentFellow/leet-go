package leetcode

func findNumbers(nums []int) int {
	ans := 0
	for _, v := range nums {
		if getSize(v)%2 == 0 {
			ans++
		}
	}

	return ans
}

func getSize(v int) int {
	ans := 0
	for v != 0 {
		ans++
		v /= 10
	}

	return ans
}
