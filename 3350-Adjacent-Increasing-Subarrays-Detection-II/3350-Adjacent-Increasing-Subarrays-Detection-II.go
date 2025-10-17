package leetcode

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)

	inc := make([]int, n)
	inc[n-1] = 1

	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			inc[i] = inc[i+1] + 1
		} else {
			inc[i] = 1
		}
	}

	ans := 0
	for i := range n - 1 {
		ans = max(ans, inc[i]/2)
		if i+inc[i] < n {
			ans = max(ans, min(inc[i], inc[i+inc[i]]))
		}
	}

	return ans
}
