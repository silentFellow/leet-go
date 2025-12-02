package leetcode

import "sort"

const MOD = 1e9 + 7

func countTrapezoids(points [][]int) int {
	hmap := make(map[int]int)
	for _, point := range points {
		y := point[1]
		hmap[y]++
	}

	ys := make([]int, 0, len(hmap))
	for y := range hmap {
		if hmap[y] >= 2 {
			ys = append(ys, y)
		}
	}
	sort.Ints(ys)

	comb := make([]int, len(ys))
	for i, c := range ys {
		cnt := hmap[c]
		comb[i] = (cnt * (cnt - 1) / 2) % MOD
	}

	pref := make([]int, len(ys)+1)
	for i := range ys {
		pref[i+1] = (pref[i] + comb[i]) % MOD
	}

	ans := 0
	for i := 0; i < len(ys); i++ {
		right := (pref[len(ys)] - pref[i+1] + MOD) % MOD
		ans = (ans + comb[i]*right) % MOD
	}

	return ans
}
