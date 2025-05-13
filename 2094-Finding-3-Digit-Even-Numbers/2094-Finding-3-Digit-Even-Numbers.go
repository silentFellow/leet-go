package leetcode

// TODO; try to use permutations for optimization

import "sort"

func findEvenNumbers(digits []int) []int {
	hmap := make(map[int]struct{})

	for i := range digits {
		if digits[i] == 0 {
			continue
		}

		for j := range digits {
			if i == j {
				continue
			}

			for k := range digits {
				if k == i || k == j {
					continue
				}

				if digits[k]%2 == 0 {
					val := (digits[i] * 100) + (digits[j] * 10) + digits[k]
					hmap[val] = struct{}{}
				}
			}
		}
	}

	val := make([]int, len(hmap))
	idx := 0

	for key := range hmap {
		val[idx] = key
		idx++
	}

	sort.Ints(val)

	return val
}
