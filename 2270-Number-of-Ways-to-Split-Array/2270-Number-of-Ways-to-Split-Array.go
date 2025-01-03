package leetcode

func waysToSplitArray(nums []int) int {
  count := 0

  totalSum := 0
  for _, num := range nums {
    totalSum += num
  }

  leftSum, rightSum := 0, totalSum
  for i:=0; i<len(nums)-1; i++ {
    num := nums[i]
    leftSum += num
    rightSum -= num

    if leftSum >= rightSum {
      count++
    }
  }

  return count
}
