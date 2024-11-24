package leetcode

import "strconv"

func reversed(val string) string {
	rune := []rune(val)

	for i, j := 0, len(rune)-1; i < j; i, j = i+1, j-1 {
		rune[i], rune[j] = rune[j], rune[i]
	}

	return string(rune)
}

func countDistinctIntegers(nums []int) int {
	var count int
	hmap := make(map[int]bool)

	for _, val := range nums {
		_, exists := hmap[val]
		if !exists {
			count++
			hmap[val] = true
		}

		reverseVal, err := strconv.Atoi(reversed(strconv.Itoa(val)))
		if err == nil {
			_, exists := hmap[reverseVal]
			if !exists {
				count++
				hmap[reverseVal] = true
			}
		}
	}

	return count
}
