package leetcode

import "slices"

func countDays(days int, meetings [][]int) int {
  slices.SortFunc(meetings, func(i, j []int) int {
    if i[0] != j[0] {
      return i[0] - j[0]
    }
    return i[1] - j[1]
  })

  prev, ans := 0, 0

  for _, meeting := range meetings {
    if prev >= meeting[0] {
      prev = max(prev, meeting[1])
      continue
    }

    ans += meeting[0] - prev - 1
    prev = meeting[1]
  }

  ans += days - prev

  return ans
}
