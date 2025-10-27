package leetcode

import (
	"math"
	"sort"
)

func nextBeautifulNumber(n int) int {
	ans := math.MaxInt

	findLen := func(arr [8]int) int {
		l := 0
		for _, v := range arr {
			l += v
		}
		return l
	}

	toValue := func(d []int) int {
		v := 0
		for _, x := range d {
			v = v*10 + x
		}
		return v
	}

	nextPerm := func(a []int) bool {
		i := len(a) - 2
		for i >= 0 && a[i] >= a[i+1] {
			i--
		}
		if i < 0 {
			return false
		}
		j := len(a) - 1
		for a[j] <= a[i] {
			j--
		}
		a[i], a[j] = a[j], a[i]
		for l, r := i+1, len(a)-1; l < r; l, r = l+1, r-1 {
			a[l], a[r] = a[r], a[l]
		}
		return true
	}

	var dfs func(arr [8]int)
	dfs = func(arr [8]int) {
		if findLen(arr) > 7 {
			return
		}

		digits := make([]int, 0, 8)
		for i := 1; i <= 7; i++ {
			for j := 0; j < arr[i]; j++ {
				digits = append(digits, i)
			}
		}

		if len(digits) > 0 {
			sort.Ints(digits)
			val := toValue(digits)
			if val > n {
				ans = min(ans, val)
			}
			for nextPerm(digits) {
				val = toValue(digits)
				if val > n {
					ans = min(ans, val)
				}
			}
		}

		for i := 1; i <= 7; i++ {
			if arr[i] == 0 {
				neoArr := arr
				neoArr[i] = i
				dfs(neoArr)
			}
		}
	}

	dfs([8]int{})
	return ans
}
