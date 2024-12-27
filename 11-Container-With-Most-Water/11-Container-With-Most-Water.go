package leetcode

func maxArea(height []int) int {
	ans := 0
	l, r := 0, len(height)-1

	for l < r {
		dis := r - l
		if height[l] < height[r] {
			ans = max(ans, height[l]*dis)
			l++
		} else {
			ans = max(ans, height[r]*dis)
			r--
		}
	}

	return ans
}
