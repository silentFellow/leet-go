package leetcode

func maximumUniqueSubarray(nums []int) int {
	hmap := make(map[int]struct{})

	ans := 0

	l := 0
	hmap[nums[0]] = struct{}{}

	cur := nums[0]
	for r := 1; r < len(nums); r++ {
		if _, ok := hmap[nums[r]]; ok {
			ans = max(ans, cur)
			for {
				if _, ok := hmap[nums[r]]; !ok {
					break
				}
				delete(hmap, nums[l])
				cur -= nums[l]
				l++
			}
		}

		hmap[nums[r]] = struct{}{}
		cur += nums[r]
	}

	ans = max(ans, cur)

	return ans
}
