package leetcode

import "math"

func maximumTripletValue(nums []int) int64 {
  n := len(nums)

  prefixMax := make([]int, len(nums))
  maxv := math.MinInt
  for i, v := range nums {
    maxv = max(maxv, v)
    prefixMax[i] = maxv
  }

  suffixMax := make([]int, len(nums))
  maxv = math.MinInt
  for i:=n-1; i>=0; i-- {
    maxv = max(maxv, nums[i])
    suffixMax[i] = maxv
  }

  ans := 0
  for i:=1; i<n-1; i++ {
    ans = max(ans, (prefixMax[i-1] - nums[i]) * suffixMax[i+1])
  }

  return int64(ans)
}
