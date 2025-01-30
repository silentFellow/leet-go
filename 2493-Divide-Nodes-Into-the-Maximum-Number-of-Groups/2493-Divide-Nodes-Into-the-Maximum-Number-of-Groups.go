package leetcode

func magnificentSets(n int, edges [][]int) int {
	graph := make([][]int, n)

	// Build adjacency list
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1 // Convert to 0-based index
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	heights := make(map[int]int) // Store max height for each component

	// Process each node, ensuring BFS is only done per connected component
	for start := 0; start < n; start++ {
		if componentID, height, ok := bfsCheckAndFindHeight(start, graph); ok {
			if h, exists := heights[componentID]; exists {
				heights[componentID] = max(h, height)
			} else {
				heights[componentID] = height
			}
		} else {
			return -1 // If any component is not bipartite, return -1
		}
	}

	// Sum up the maximum heights of all connected components
	totalDiameters := 0
	for _, height := range heights {
		totalDiameters += height
	}
	return totalDiameters
}

// bfsCheckAndFindHeight performs BFS from `start`, verifies bipartiteness, and finds max height in component.
func bfsCheckAndFindHeight(start int, graph [][]int) (componentID int, maxHeight int, ok bool) {
	level := make([]int, len(graph)) // Track BFS levels
	queue := []int{start}
	level[start] = 1
	maxHeight = 1
	componentID = start // Smallest node in component

	for front := 0; front < len(queue); front++ {
		v := queue[front]
		maxHeight = level[v] // Update max BFS level
		componentID = min(componentID, v)

		// Traverse neighbors and check bipartiteness
		for _, u := range graph[v] {
			if level[u] == 0 { // First visit to `u`
				level[u] = level[v] + 1
				queue = append(queue, u)
			} else if absDiff(level[u], level[v]) != 1 { // Check bipartiteness
				ok = false
				return
			}
		}
	}

	ok = true
	return
}

// absDiff computes the absolute difference between two integers.
func absDiff(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}
