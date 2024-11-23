package leetcode

import "math"

func findDuplicate(nums []int) int {
  var repeated int
  for _, val := range nums {
    val = int(math.Abs(float64(val)))
    if nums[val-1] < 0 {
      return val
    }

    nums[val-1] = -nums[val-1]
  }

  return repeated
}
