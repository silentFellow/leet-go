package leetcode

import "math"

type pos struct {
  r, c int
}

func minimumTotal(triangle [][]int) int {
  n := len(triangle)
  hmap := make(map[pos]int)

  var dfs func(r, c int) int
  dfs = func(r, c int) int {
    if val, ok := hmap[pos{r, c}]; ok {
      return val
    } else if r == n-1 {
      return triangle[r][c]
    } else if c >= len(triangle[r]) {
      return math.MaxInt64
    }

    first := dfs(r+1, c)
    second := dfs(r+1, c+1)

    val := triangle[r][c] + min(first, second)
    hmap[pos{r, c}] = val
    return val
  }

  return dfs(0, 0)
}
