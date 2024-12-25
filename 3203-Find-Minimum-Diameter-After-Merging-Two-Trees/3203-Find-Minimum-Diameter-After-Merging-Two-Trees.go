package leetcode

import (
	"math"
	"sort"
)

type adjList map[int][]int

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
  adj1, adj2 := build_adj(edges1), build_adj(edges2)

  var dfs func(current, parent int, list adjList) (int, int)
  dfs = func(c, p int, l adjList) (int, int) {
    max_d := 0
    max_child_pairs := [2]int{0, 0}

    for _, neighbour := range l[c] {
      if neighbour == p { continue }
      _, nei_max_child_pairs := dfs(neighbour, c, l)

      if nei_max_child_pairs > max_child_pairs[0] {
        max_child_pairs[0] = nei_max_child_pairs
        sort.Ints(max_child_pairs[:])
      }
    }

    max_d = max(max_d, max_child_pairs[0] + max_child_pairs[1])
    return max_d, (1 + max(max_child_pairs[0], max_child_pairs[1]))
  }

  d1, _ := dfs(0, -1, adj1)
  d2, _ := dfs(0, -1, adj2)

  max3 := func(f, s, t int) int {
    return max(f, max(s, t))
  }

  return max3(d1, d2, 1 + int(math.Ceil(float64(d1) / 2)) + int(math.Ceil(float64(d2) / 2)))
}

func build_adj(edges [][]int) adjList {
  hmap := make(adjList)

  for _, val := range edges {
    hmap[val[0]] = append(hmap[val[0]], val[1])
    hmap[val[1]] = append(hmap[val[1]], val[0])
  }

  return hmap
}
