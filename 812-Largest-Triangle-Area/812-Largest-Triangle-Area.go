package leetcode

func largestTriangleArea(points [][]int) float64 {
	n := len(points)

	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	// Shoelace's Formula (also known as Gauss's area formula and surveyor's formula):
	// Area = (1/2) * |x1 * (y2 - y3) + x2 * (y3 - y1) + x3 * (y1 - y2)|
	// order important
	area := func(a, b, c []int) float64 {
		return 0.5 * float64(abs(a[0]*(b[1]-c[1])+b[0]*(c[1]-a[1])+c[0]*(a[1]-b[1])))
	}

	var ans float64
	for i := range n {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				ans = max(ans, area(points[i], points[j], points[k]))
			}
		}
	}

	return ans
}
