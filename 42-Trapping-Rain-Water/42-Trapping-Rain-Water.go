package leetcode

func trap(height []int) int {
	l, r := 0, len(height)-1
	lh, rh := height[l], height[r]

	ans := 0

	for l < r {
		if lh < rh {
			l++
			lh = max(lh, height[l])
			ans += lh - height[l]
		} else {
			r--
			rh = max(rh, height[r])
			ans += rh - height[r]
		}
	}

	return ans
}
