package leetcode

func check(nums []int) bool {
  n := len(nums)
	once := false

	for i := 0; i < n; i++ {
		if nums[i] < nums[(i-1+n)%n] {
			if once {
				return false
			}
			once = true
		}
	}

	return true
}
