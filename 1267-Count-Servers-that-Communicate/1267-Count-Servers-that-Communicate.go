package leetcode

func countServers(grid [][]int) int {
  rowCount, colCount := make([]int, len(grid)), make([]int, len(grid[0]))

  for i, row := range grid {
    for j, col := range row {
      if col == 1 {
        rowCount[i]++
        colCount[j]++
      }
    }
  }

  ans := 0
  for i, row := range grid {
    for j, col := range row {
      if col == 1 && (rowCount[i] > 1 || colCount[j] > 1) {
        ans++
      }
    }
  }

  return ans
}
