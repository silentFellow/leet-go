package leetcode

func largestRectangleArea(heights []int) int {
	ans := 0

	// to stimulate force cleanup
	heights = append(heights, -1)

	stack := []int{}
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			n := len(stack)
			topIdx := stack[n-1]
			stack = stack[:n-1]

			height := heights[topIdx]
			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			ans = max(ans, height*width)
		}
		stack = append(stack, i)
	}

	return ans
}
