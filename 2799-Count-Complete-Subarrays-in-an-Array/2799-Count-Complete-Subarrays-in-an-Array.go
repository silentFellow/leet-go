package leetcode

func countCompleteSubarrays(nums []int) int {
	n := len(nums)

	hmap := make(map[int]struct{})
	for _, v := range nums {
		hmap[v] = struct{}{}
	}

	req := len(hmap)
	newMap := make(map[int]int)

	ans := 0
	left := 0

	for right := range n {
		newMap[nums[right]]++

		for len(newMap) == req {
			ans += n - right

			newMap[nums[left]]--
			if newMap[nums[left]] == 0 {
				delete(newMap, nums[left])
			}
			left++
		}
	}

	return ans
}
