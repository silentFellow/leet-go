package leetcode

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	adjList := buildAdj(prerequisites)

	hmap := make(map[int]map[int]struct{})
	var dfs func(cur int) map[int]struct{}
	dfs = func(cur int) map[int]struct{} {
		if val, ok := hmap[cur]; ok {
			return val
		}

		prereqMap := make(map[int]struct{})
		for _, prereq := range adjList[cur] {
			prereqMap = findUnion(prereqMap, dfs(prereq))
		}
		prereqMap[cur] = struct{}{}
		hmap[cur] = prereqMap

    return prereqMap
	}

	for i := range numCourses {
		dfs(i)
	}

	result := make([]bool, len(queries))
	for i, query := range queries {
		map1, ok := hmap[query[1]]
    _, ok = map1[query[0]]
		result[i] = ok
	}

	return result
}

func buildAdj(graph [][]int) map[int][]int {
	hmap := make(map[int][]int)

	for _, node := range graph {
		hmap[node[1]] = append(hmap[node[1]], node[0])
	}

	return hmap
}

func findUnion(u, v map[int]struct{}) map[int]struct{} {
  hmap := make(map[int]struct{})

  for key, val := range u {
    hmap[key] = val
  }

  for key, val := range v {
    hmap[key] = val
  }

  return hmap
}
