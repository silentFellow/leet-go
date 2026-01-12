package leetcode

func minTimeToVisitAllPoints(points [][]int) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	movesReq := func(x1, x2 []int) int {
		f, s := abs(x1[0]-x2[0]), abs(x1[1]-x2[1])
		return max(f, s)
	}

	ans := 0
	for i := range len(points) - 1 {
		ans += movesReq(points[i], points[i+1])
	}

	return ans
}
