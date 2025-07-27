package leetcode

func countHillValley(nums []int) int {
	cleaned := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			cleaned = append(cleaned, nums[i])
		}
	}

	ans := 0
	for i := 1; i < len(cleaned)-1; i++ {
		f, s, t := cleaned[i-1], cleaned[i], cleaned[i+1]

		if s < f && s < t {
			ans++
		} else if s > f && s > t {
			ans++
		}
	}

	return ans
}
