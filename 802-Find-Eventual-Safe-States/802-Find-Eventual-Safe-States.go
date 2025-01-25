package leetcode

func eventualSafeNodes(graph [][]int) []int {
  state := make([]int, len(graph)) // 0 => unvisited, 1 => visiting, 2 => safe

  var dfs func(v int) bool
  dfs = func(v int) bool {
    if state[v] > 0 {
      return state[v] == 2
    }
    state[v] = 1

    for _, nei := range graph[v] {
      if !dfs(nei) {
        return false
      }
    }

    state[v] = 2
    return true
  }

  ans := []int{}
  for i := range graph {
    if dfs(i) {
      ans = append(ans, i)
    }
  }

  return ans
}
