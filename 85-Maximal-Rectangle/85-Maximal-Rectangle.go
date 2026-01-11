package leetcode

func maximalRectangle(matrix [][]byte) int {
	maxArea := func(heights []int) int {
		area := 0

		// force cleanup
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
				area = max(area, height*width)
			}
			stack = append(stack, i)
		}

		return area
	}

	ans := 0
	cur := make([]int, len(matrix[0]))
	for _, row := range matrix {
		for i, col := range row {
			if col == '0' {
				cur[i] = 0
			} else {
				cur[i]++
			}
		}

		ans = max(ans, maxArea(cur))
	}

	return ans
}
