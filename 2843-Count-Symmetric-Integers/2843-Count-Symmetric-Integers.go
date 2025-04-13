package leetcode

import "strconv"

func sum(v int) int {
	ans := 0

	for v != 0 {
		ans += v % 10
		v /= 10
	}

	return ans
}

func countSymmetricIntegers(low int, high int) int {
	ans := 0

	for i := low; i < high+1; i++ {
		str := strconv.Itoa(i)
		n := len(str)

		if n%2 == 1 {
			continue
		}

		mid := n / 2
		fv, fe := strconv.Atoi(str[:mid])
		sv, se := strconv.Atoi(str[mid:])
		if fe != nil || se != nil {
			continue
		}

		if sum(fv) == sum(sv) {
			ans++
		}
	}

	return ans
}
