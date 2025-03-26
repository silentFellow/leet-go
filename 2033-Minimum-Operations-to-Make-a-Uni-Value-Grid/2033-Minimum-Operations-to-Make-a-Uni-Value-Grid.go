package leetcode

import "slices"

func minOperations(grid [][]int, x int) int {
  uni, idx := make([]int, len(grid)*len(grid[0])), 0
  for _, row := range grid {
    for _, col := range row {
      uni[idx] = col
      idx++
    }
  }

  first := uni[0] % x
  for _, v := range uni {
    if v % x != first {
      return -1
    }
  }

  slices.Sort(uni)
  req := uni[len(uni)/2]
  ans := 0
  for _, v := range uni {
    ans += abs(v-req)/x
  }

  return ans
}

func abs(v int) int {
  if v < 0 {
    return -v
  }

  return v
}
