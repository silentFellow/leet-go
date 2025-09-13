package leetcode

func countPairs(nums []int, target int) int {
	n := len(nums)

	ans := 0
	for i := range n {
		for j:=i+1; j<n; j++ {
			if nums[i] + nums[j] < target {
				ans++
			}
		}
	}

	return ans
}
