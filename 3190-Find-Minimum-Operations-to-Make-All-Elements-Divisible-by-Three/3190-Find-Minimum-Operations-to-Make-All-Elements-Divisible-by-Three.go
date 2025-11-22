package leetcode

func minimumOperations(nums []int) int {
	ans := 0

	for _, v := range nums {
		if v % 3 == 0 {
			continue
		} else if ((v+1) % 3 == 0) || ((v-1) % 3 == 0) {
			ans++
		} else {
			ans += 2
		}
	}

	return ans
}
