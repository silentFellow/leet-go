package leetcode

func maximumLength(nums []int) int {
	allOdd, allEven := 0, 0
	for _, v := range nums {
		if v%2 == 0 {
			allEven++
		} else {
			allOdd++
		}
	}

	altOdd, prevIsOdd := 0, false
	for _, v := range nums {
		if prevIsOdd && v%2 == 0 {
			altOdd++
			prevIsOdd = false
		} else if !prevIsOdd && v%2 == 1 {
			altOdd++
			prevIsOdd = true
		}
	}

	alteven, prevIsEven := 0, false
	for _, v := range nums {
		if prevIsEven && v%2 == 1 {
			alteven++
			prevIsEven = false
		} else if !prevIsEven && v%2 == 0 {
			alteven++
			prevIsEven = true
		}
	}

	return max(allEven, allOdd, altOdd, alteven)
}
