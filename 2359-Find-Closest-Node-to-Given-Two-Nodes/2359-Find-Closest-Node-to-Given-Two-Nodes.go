package leetcode

import "math"

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	getDistance := func(start int) []int {
		dist := make([]int, len(edges))
		for i := range len(edges) {
			dist[i] = -1
		}

		d := 0
		for start != -1 && dist[start] == -1 {
			dist[start] = d
			d++
			start = edges[start]
		}

		return dist
	}

	dist1, dist2 := getDistance(node1), getDistance(node2)
	ans := -1
	minDistance := math.MaxInt

	for i := range len(edges) {
		if dist1[i] != -1 && dist2[i] != -1 {
			maxv := max(dist1[i], dist2[i])
			if maxv < minDistance {
				minDistance = maxv
				ans = i
			}
		}
	}

	return ans
}
