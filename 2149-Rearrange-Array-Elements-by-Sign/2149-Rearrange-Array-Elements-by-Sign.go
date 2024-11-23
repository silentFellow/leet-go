package leetcode

func rearrangeArray(nums []int) []int {
  pos := make([]int, 0, len(nums)/2)
  neg := make([]int, 0, len(nums)/2)

  for _, val := range nums {
    if val > 0 {
      pos = append(pos, val)
    } else {
      neg = append(neg, val)
    }
  }

  index := 0
  for i := range len(pos) {
    nums[index] = pos[i]
    nums[index+1] = neg[i]
    index += 2
  }

  return nums
}
