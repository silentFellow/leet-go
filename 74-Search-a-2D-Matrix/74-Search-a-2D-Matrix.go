package leetcode

func searchMatrix(matrix [][]int, target int) bool {
  l, r := 0, len(matrix[0])-1

  for l < len(matrix) && r >= 0 {
    if matrix[l][r] == target {
      return true
    }

    if matrix[l][r] > target {
      r--
    } else {
      l++
    }
  }

  return false
}
