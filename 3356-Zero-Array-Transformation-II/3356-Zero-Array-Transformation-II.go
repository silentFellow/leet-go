package leetcode

// TODO: will get TLE optimize it

func minZeroArray(nums []int, queries [][]int) int {
  zeroCount, ans := 0, -1

  for _, v := range nums {
    if v == 0 {
      zeroCount++
    }
  }
  if zeroCount == len(nums) {
    return 0
  }

  for idx, query := range queries {
    l, r, v := query[0], query[1], query[2]

    for i:=l; i<=r; i++ {
      if nums[i] == 0 {
        continue
      }

      if nums[i] - v <= 0 {
        nums[i] = 0
        zeroCount++
      } else {
        nums[i] -= v
      }
    }

    if zeroCount == len(nums) {
      ans = idx+1
      break
    }
  }

  return ans
}
