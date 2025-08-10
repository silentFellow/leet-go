package leetcode

func reorderedPowerOf2(n int) bool {
	found := false

	narr := make([]int, 0)
	for n != 0 {
		narr = append(narr, n%10)
		n /= 10
	}

	join := func(arr []int) int {
		joined := 0
		for _, v := range arr {
			joined = (joined * 10) + v
		}
		return joined
	}

	var permutate func(arr []int, l, r int)
	permutate = func(arr []int, l, r int) {
		if found {
			return
		}

		if l == r {
			if arr[0] != 0 {
				val := join(arr)
				if val > 0 && (val&(val-1) == 0) {
					found = true
				}
			}

			return
		}

		for i := l; i <= r; i++ {
			arr[i], arr[l] = arr[l], arr[i]
			permutate(arr, l+1, r)
			arr[i], arr[l] = arr[l], arr[i]
		}
	}

	permutate(narr, 0, len(narr)-1)
	return found
}
