package leetcode

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		rootX := find(x)
		rootY := find(y)
		if rootX == rootY {
			return false
		}
		parent[rootX] = rootY
		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}

	return nil
}
