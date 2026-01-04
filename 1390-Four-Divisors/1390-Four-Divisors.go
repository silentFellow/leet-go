package leetcode

func sumFourDivisors(nums []int) int {
	// number => no. of divisors, sum of divisors, no. of keys(to find duplicates)
	hmap := make(map[int][3]int)

	for _, num := range nums {
		if _, ok := hmap[num]; !ok {
			hmap[num] = [3]int{0, 0, 1}
		} else { // duplicates
			val := hmap[num]
			val[2]++
			hmap[num] = val
			continue
		}

		val := hmap[num]

		for i := 1; i <= num; i++ {
			if num%i == 0 {
				val[0]++
				val[1] += i

				if val[0] > 4 {
					break
				}
			}
		}
		hmap[num] = val
	}

	ans := 0
	for _, v := range hmap {
		n, sum, cnt := v[0], v[1], v[2]
		// if same key repeated twice
		if n == 4 {
			ans += (sum * cnt)
		}
	}

	return ans
}
