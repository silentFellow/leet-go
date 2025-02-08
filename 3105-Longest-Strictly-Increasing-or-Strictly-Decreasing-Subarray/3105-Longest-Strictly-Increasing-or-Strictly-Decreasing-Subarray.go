package leetcode

func longestMonotonicSubarray(nums []int) int {
  n := len(nums)
  if n < 2 {
    return n
  }

  inc, dec, ans := 1, 1, 0
  for i:=1; i<n; i++ {
    if nums[i] < nums[i-1] {
      dec++
      inc = 1
    }

    if nums[i] > nums[i-1] {
      inc++
      dec = 1
    }

    if nums[i] == nums[i-1] {
      inc, dec = 1, 1
    }

    ans = max(ans, inc, dec)
  }

  return ans
}
