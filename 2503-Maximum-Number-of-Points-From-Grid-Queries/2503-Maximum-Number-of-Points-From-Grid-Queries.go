package leetcode

// TODO: optimize to remove tle

import "fmt"

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])
	ans := make([]int, len(queries))
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var bfs func(thresh int) int
	bfs = func(thresh int) int {
		if grid[0][0] >= thresh {
			return 0
		}

		visited := make(map[string]struct{})
		visited[fmt.Sprintf("%v-%v", 0, 0)] = struct{}{}
		queue := [][]int{{0, 0}}
		time := 1

		for len(queue) > 0 {
			oldLen := len(queue)

			for _, v := range queue {
				for _, dir := range directions {
					rm, cm := (v[0] + dir[0]), (v[1] + dir[1])
					if rm >= 0 && rm < m && cm >= 0 && cm < n && grid[rm][cm] < thresh {
						if _, ok := visited[fmt.Sprintf("%v-%v", rm, cm)]; !ok {
							queue = append(queue, []int{rm, cm})
              visited[fmt.Sprintf("%v-%v", rm, cm)] = struct{}{}
							time++
						}
					}
				}
			}

			queue = queue[oldLen:]
		}

		return time
	}

	for i, query := range queries {
		ans[i] = bfs(query)
	}

	return ans
}
