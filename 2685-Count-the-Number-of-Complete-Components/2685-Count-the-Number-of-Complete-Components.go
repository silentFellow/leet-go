package leetcode

func countCompleteComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	for i := range n {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}

		return parent[x]
	}

	var union func(x, y int)
	union = func(x, y int) {
		parent[find(y)] = find(x)
	}

	for _, v := range edges {
		union(v[0], v[1])
	}

	rootCounts := make(map[int][]int)
	for i := range n {
		root := find(i)
		rootCounts[root] = append(rootCounts[root], i)
	}

	edgeCount := make(map[int]int)
	for _, edge := range edges {
		root := find(edge[0])
		edgeCount[root]++
	}

	completeNodes := 0
	for root, nodes := range rootCounts {
		n := len(nodes)
		expected := n * (n-1) / 2
		if edgeCount[root] == expected {
			completeNodes++
		}
	}

	return completeNodes
}
