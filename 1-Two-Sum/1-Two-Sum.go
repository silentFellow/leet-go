package leetcode

func twoSum(nums []int, target int) []int {
  hmap := make(map[int]int)

  for i, val := range nums {
    req := target-val
    existIndex, exists := hmap[req]
    if exists {
      return []int{ existIndex, i }
    }

    hmap[val] = i
  }

  return []int{-1, -1}
}
