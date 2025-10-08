package leetcode

import (
	"slices"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)

	ans := make([]int, len(spells))
	for i, spell := range spells {
		start := sort.Search(len(potions), func(potionIdx int) bool {
			return int64(spell)*int64(potions[potionIdx]) >= success
		})

		ans[i] = len(potions) - start
	}

	return ans
}
