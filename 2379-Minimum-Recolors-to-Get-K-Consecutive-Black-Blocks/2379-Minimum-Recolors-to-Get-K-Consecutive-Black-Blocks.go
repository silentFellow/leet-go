package leetcode

import "math"

func minimumRecolors(blocks string, k int) int {
  ans := math.MaxInt

  b := 0
  for i := range k {
    if blocks[i] == 'B' {
      b++
    }
  }
  ans = min(ans, returnModified(k, b))

  for i:=1; i<=len(blocks)-k; i++ {
    if blocks[i-1] == 'B' {
      b--
    }

    if blocks[i+k-1] == 'B' {
      b++
    }

    ans = min(ans, returnModified(k, b))
  }

  return ans
}

func returnModified(k, v int) int {
  if k-v < 0 {
    return 0
  }

  return k-v
}
