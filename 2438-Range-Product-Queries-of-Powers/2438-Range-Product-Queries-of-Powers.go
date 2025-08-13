package leetcode

const MOD = 1_000_000_007

func modPow(x, y int) int {
	res := 1
	base := x % MOD
	exp := y
	for exp > 0 {
		if exp&1 == 1 {
			res = (res * base) % MOD
		}
		base = (base * base) % MOD
		exp >>= 1
	}
	return res
}

func modInverse(x int) int {
	return modPow(x, MOD-2)
}

func productQueries(n int, queries [][]int) []int {
	getBinary := func(v int) []int {
		res := make([]int, 0)
		for v != 0 {
			res = append(res, v%2)
			v /= 2
		}

		return res
	}

	binArr := getBinary(n)
	arr := make([]int, 0)
	for i, v := range binArr {
		if v != 0 {
			val := 1 << i
			arr = append(arr, val%MOD)
		}
	}

	preProcess := make([]int, len(arr)+1)
	preProcess[0] = 1
	for i := range len(arr) {
		preProcess[i+1] = (preProcess[i] * arr[i]) % MOD
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		f, s := query[0], query[1]
		ans[i] = (preProcess[s+1] * modInverse(preProcess[f])) % MOD
	}

	return ans
}
