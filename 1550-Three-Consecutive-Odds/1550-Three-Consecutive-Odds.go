package leetcode

func threeConsecutiveOdds(arr []int) bool {
	oddCount := 0

	for _, v := range arr {
		if v%2 == 0 {
			oddCount = 0
		} else {
			oddCount++
		}

		if oddCount == 3 {
			return true
		}
	}

	return false
}
