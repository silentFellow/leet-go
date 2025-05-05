package leetcode

func numTilings(n int) int {
	MOD := 1000000007

	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	f, s, t := 1, 1, 2

	for i := 3; i <= n; i++ {
		f, s, t = s, t, (t*2+f)%MOD
	}

	return t
}

// Array-based DP
// ans := make([]int, n+2)
// ans[0], ans[1], ans[2] = 1, 1, 2
//
// if n <= 2 {
// 	return ans[n]
// }
//
// for i := 3; i <= n; i++ {
// 	ans[i] = (ans[i-1]*2 + ans[i-3]) % MOD
// }
//
// return ans[n]
