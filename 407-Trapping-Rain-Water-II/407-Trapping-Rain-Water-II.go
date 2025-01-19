package leetcode

import (
	"container/heap"
)

type Cell struct {
	Height, Row, Col int
}

type MinHeap []Cell

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i].Height < h[j].Height }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Cell)) }

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func trapRainWater(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}

	h := &MinHeap{}
	heap.Init(h)
	maxCurrHeight := 0
	volume := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || i == n-1 || j == 0 || j == m-1 {
				heap.Push(h, Cell{grid[i][j], i, j})
				vis[i][j] = true
			}
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for h.Len() > 0 {
		cell := heap.Pop(h).(Cell)
		maxCurrHeight = max(maxCurrHeight, cell.Height)

		for _, d := range directions {
			nx, ny := cell.Row+d[0], cell.Col+d[1]
			if nx >= 0 && ny >= 0 && nx < n && ny < m && !vis[nx][ny] {
				vis[nx][ny] = true
				if grid[nx][ny] < maxCurrHeight {
					volume += maxCurrHeight - grid[nx][ny]
				}
				heap.Push(h, Cell{grid[nx][ny], nx, ny})
			}
		}
	}

	return volume
}
