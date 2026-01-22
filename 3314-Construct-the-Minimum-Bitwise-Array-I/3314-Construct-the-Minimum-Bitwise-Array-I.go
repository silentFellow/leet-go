package leetcode

func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i, num := range nums {
		if num%2 == 0 {
			ans[i] = -1
		} else {
			for j := 0; j <= num; j++ {
				if j|(j+1) == num {
					ans[i] = j
					break
				}
			}
		}
	}

	return ans
}
