package leetcode

func maxCount(banned []int, n int, maxSum int) int {
  hmap := make(map[int]struct{})
  for _, val := range banned {
    hmap[val] = struct{}{}
  }

  count, sum := 0, 0

  for i:=1; i<=n; i++ {
    if _, ok := hmap[i]; !ok {
      if sum + i > maxSum {
        break
      }

      count++
      sum += i
    }
  }

  return count
}
