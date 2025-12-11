package leetcode

const MOD = 1e9 + 7

func countPermutations(complexity []int) int {
	minv := complexity[0]

	freqMap := make(map[int]int)
	for _, c := range complexity {
		freqMap[c]++

		if c < minv {
			return 0
		}
	}

	if freqMap[minv] > 1 {
		return 0
	}

	ans := 1
	for i := 1; i < len(complexity); i++ {
		ans = (ans * i) % MOD
	}

	return ans
}
