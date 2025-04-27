package leetcode

func countSubarrays(nums []int) int {
	ans := 0

	for i := 2; i < len(nums); i++ {
		if ((nums[i] + nums[i-2]) == (nums[i-1] / 2)) && (nums[i-1]%2 == 0) {
			ans++
		}
	}

	return ans
}
