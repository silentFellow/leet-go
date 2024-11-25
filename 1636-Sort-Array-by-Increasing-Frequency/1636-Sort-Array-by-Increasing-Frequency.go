package leetcode

import "sort"

type pair struct {
  val int
  count int
}

func frequencySort(nums []int) []int {
  hmap := make(map[int]int)
  for _, val := range nums {
    hmap[val] = hmap[val] + 1
  }

  freqPairs := make([]pair, len(hmap))
  for val, count := range hmap {
    freqPairs = append(freqPairs, pair{val, count})
  }

  sort.Slice(freqPairs, func(i, j int) bool {
    if freqPairs[i].count == freqPairs[j].count {
      return freqPairs[i].val > freqPairs[j].val
    }
    return freqPairs[i].count < freqPairs[j].count
  })

  i := 0
  for _, freq := range freqPairs {
    for range freq.count {
      nums[i] = freq.val
      i++
    }
  }

  return nums
}
