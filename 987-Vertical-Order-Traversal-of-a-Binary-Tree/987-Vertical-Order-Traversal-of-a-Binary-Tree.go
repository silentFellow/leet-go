package leetcode

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeTraversal struct {
	node  *TreeNode
	line  int
	depth int
}

func verticalTraversal(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	lineMapping := make(map[int][][2]int)
	queue := []TreeNodeTraversal{}

	queue = append(queue, TreeNodeTraversal{node: root, line: 0, depth: 0})

	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]

		lineMapping[elem.line] = append(lineMapping[elem.line], [2]int{elem.depth, elem.node.Val})

		if elem.node.Left != nil {
			queue = append(
				queue,
				TreeNodeTraversal{
					node:  elem.node.Left,
					line:  elem.line - 1,
					depth: elem.depth + 1,
				},
			)
		}

		if elem.node.Right != nil {
			queue = append(
				queue,
				TreeNodeTraversal{
					node:  elem.node.Right,
					line:  elem.line + 1,
					depth: elem.depth + 1,
				},
			)
		}
	}

	keys := []int{}
	for key := range lineMapping {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, key := range keys {
		nodes := lineMapping[key]
		sort.Slice(nodes, func(i, j int) bool {
			if nodes[i][0] == nodes[j][0] {
				return nodes[i][1] < nodes[j][1]
			}

			return nodes[i][0] < nodes[j][0]
		})

		columns := []int{}
		for _, node := range nodes {
			columns = append(columns, node[1])
		}
		ans = append(ans, columns)
	}

	return ans
}
