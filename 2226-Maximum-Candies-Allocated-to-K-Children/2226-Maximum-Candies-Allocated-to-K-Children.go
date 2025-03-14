package leetcode

func maximumCandies(candies []int, k int64) int {
	ans := 0
	l, r := 1, findMax(candies)

	for l <= r {
		m := (l + r) / 2
		possible := 0

		for _, candy := range candies {
			possible += candy / m
		}

		if int64(possible) >= k {
			ans = max(ans, m)
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return ans
}

func findMax(arr []int) int {
  maxv := arr[0]

  for _, v := range arr {
    maxv = max(maxv, v)
  }

  return maxv
}
