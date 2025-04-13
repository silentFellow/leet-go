package leetcode

import "math"

var MOD = int(math.Pow(10, 9) + 7)

func countGoodNumbers(n int64) int {
	// brute force
	// ans := 1
	// for range even {
	//   ans *= 5
	//   ans %= MOD
	// }
	//
	// for range prime {
	//   ans *= 4
	//   ans %= MOD
	// }

	even := (n / 2) + (n % 2)
	prime := n / 2

	ans := (modPow(5, int(even)) * modPow(4, int(prime))) % MOD

	return int(ans)
}

func modPow(base, exp int) int {
	res := 1

	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % MOD
		}
		base = (base * base) % MOD
		exp /= 2
	}

	return res
}
