package leetcode

func findClosest(x int, y int, z int) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	f, s := abs(x-z), abs(y-z)

	if f == s {
		return 0
	} else if f < s {
		return 1
	}

	return 2
}
