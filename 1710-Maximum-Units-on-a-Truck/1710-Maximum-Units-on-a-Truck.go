package leetcode

import "slices"

func maximumUnits(boxTypes [][]int, truckSize int) int {
  slices.SortFunc(boxTypes, func(i, j []int) int {
    return -(i[1] - j[1])
  })

  ans := 0
  for _, box := range boxTypes {
    if box[0] <= truckSize {
      ans += (box[0] * box[1])
      truckSize -= box[0]
    } else {
      ans += (truckSize * box[1])
      break
    }
  }

  return ans
}
