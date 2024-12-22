package leetcode

// INTITUTION: sum of all nodes in the tree is divisible by k
// if a certain part of the tree is divisible by k
// then addition of the certain part to any other part changes nothinig

// Eg: values=(100, 3, 5, 2), k=10
// if 100 divisible by 10
// then addition of that 100 to other like:
// 100 + 3 == (103 % 10) != 0
// 103 + 5 == (108 % 10) != 0
// 108 + 2 == (110 % 10) != 0

// by this, we can prove that:
// addition of already divided part to anything no make change
// a entirely new part is required

// custom type to hold connectted edges
type adjMatrix map[int][]int

func maxKDivisibleComponents(_ int, edges [][]int, values []int, k int) int {
	adj := makeAdjMatrix(edges) // create a adjacency matrix

	res := 0 // ansert

  // bottom-up dfs function
  // that take current and parent
  // since, it contains undirected edges
  // to avooid infinite loop check if neighbour is parent
  // adds res, if that part of the component is divisible by 3
	var dfs func(current, parent int) int
	dfs = func(current, parent int) int {
		sum := values[current]

		for _, neighbours := range adj[current] {
			if neighbours != parent {
				sum += dfs(neighbours, current)
			}
		}

		if sum%k == 0 {
			res++
		}

		return sum
	}

	dfs(0, -1)

	return res
}

func makeAdjMatrix(edges [][]int) adjMatrix {
	adj := make(adjMatrix)

	for _, val := range edges {
		v1 := val[0]
		v2 := val[1]

		adj[v1] = append(adj[v1], v2)
		adj[v2] = append(adj[v2], v1)
	}

	return adj
}
