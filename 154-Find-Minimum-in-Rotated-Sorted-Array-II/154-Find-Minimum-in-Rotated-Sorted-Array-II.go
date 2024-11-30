package leetcode

func findMin(nums []int) int {
  l, r := 0, len(nums)-1

  for l < r {
    m := (l+r)/2

    // if all 3 value are equal cannot find which side is sorted
    if nums[l] == nums[m] && nums[m] == nums[r] {
      l++
      r--
      continue
    }

    // left sort
    if nums[m] > nums[r] {
      l = m+1
    } else {
      r = m
    }
  }

  return nums[l]
}
