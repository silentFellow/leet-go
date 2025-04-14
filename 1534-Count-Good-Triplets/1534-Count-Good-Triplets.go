package leetcode

// TODO: optimize it

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n := len(arr)
	ans := 0

	for i, v := range arr {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				f, s, t := abs(v-arr[j]), abs(arr[j]-arr[k]), abs(v-arr[k])
				if f <= a && s <= b && t <= c {
					ans++
				}
			}
		}
	}

	return ans
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
