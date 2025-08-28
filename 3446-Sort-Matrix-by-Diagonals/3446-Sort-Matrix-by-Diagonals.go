package leetcode

import "sort"

func sortMatrix(grid [][]int) [][]int {
	hmap := make(map[int][]int)

	for i, row := range grid {
		for j, col := range row {
			v := i - j
			hmap[v] = append(hmap[v], col)
		}
	}

	for k, v := range hmap {
		if k >= 0 {
			sort.Sort(sort.Reverse(sort.IntSlice(v)))
		} else {
			sort.Ints(v)
		}
		hmap[k] = v
	}

	for i, row := range grid {
		for j := range row {
			v := i - j

			grid[i][j] = hmap[v][0]
			hmap[v] = hmap[v][1:]
		}
	}

	return grid
}
