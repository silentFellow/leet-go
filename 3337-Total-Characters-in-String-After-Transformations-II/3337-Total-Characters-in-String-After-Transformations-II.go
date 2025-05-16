package leetcode

func lengthAfterTransformations(s string, t int, nums []int) int {
	// MODULO
	const MOD = 1_000_000_007

	// initial array
	cur := [26]int{}
	for _, v := range s {
		cur[v-'a']++
	}

	// calculation
	for range t {
		temp := [26]int{}

		for ch := range 26 {
			times := cur[ch]
			limit := nums[ch]

			for j := 1; j <= limit; j++ {
				idx := (ch + j) % 26
				temp[idx] = (temp[idx] + times) % MOD
			}
		}
		cur = temp
	}

	// answer calculation
	ans := 0
	for _, v := range cur {
		ans = (ans + v) % MOD
	}

	return ans
}
