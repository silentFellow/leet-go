package leetcode

func zeroFilledSubarray(nums []int) int64 {
	ans := 0

	zc := 0
	for _, v := range nums {
		if v == 0 {
			zc++
		} else {
			if zc != 0 {
				ans += (zc * (zc + 1)) / 2
				zc = 0
			}
		}
	}

	if zc != 0 {
		ans += (zc * (zc + 1)) / 2
		zc = 0
	}

	return int64(ans)
}
