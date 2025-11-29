package leetcode

const MOD = int(1e9 + 7)

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	type key struct{ i, j, rem int }
	hmap := make(map[key]int)

	var dfs func(i, j, rem int) int
	dfs = func(i, j, rem int) int {
		if i >= m || j >= n {
			return 0
		}

		rem = (rem + grid[i][j]) % k
		if i == m-1 && j == n-1 {
			if rem == 0 {
				return 1
			}
			return 0
		}

		k := key{i, j, rem}
		if v, ok := hmap[k]; ok {
			return v
		}

		res := (dfs(i, j+1, rem) + dfs(i+1, j, rem)) % MOD
		hmap[k] = res
		return res
	}

	return dfs(0, 0, 0)
}
