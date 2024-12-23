package leetcode

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) int {
	// base condition
	if root == nil {
		return 0
	}

	// aggregate result
	ans := 0

	// queue for level order traversal
	queue := []*TreeNode{root}
	for len(queue) > 0 { // till all node is visited once
		// get all values in the level in order
		values := []int{}
		for _, node := range queue {
			values = append(values, node.Val)
		}
		ans += minSwap(values) // find minimum swap and add to result

		// append next level
		n := len(queue)
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// remove already visited nodes
		queue = queue[n:]
	}

	return ans // return result
}

// minimum swap part
type query struct {
	val int
	idx int
}

type pair []query

// utility function to swap two values
// along with the paired index
func (p *pair) swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

// cyclic detection algorithm
func minSwap(nums []int) int {
	// create a pair of value with index in the nums
	var pairs pair
	for i, val := range nums {
		newPair := query{
			val: val,
			idx: i,
		}

		pairs = append(pairs, newPair)
	}

	// targeted array
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val < pairs[j].val
	})

	swaps := 0 // counter
	for i := 0; i < len(nums); i++ {
		if pairs[i].idx == i { // if in correct position
			continue
		} else {
			// if not in correct position
      // increment counter
      // swap the array element, both value and index

      // decrease i, so that it stays in the current index in next loop
      // to check whether the swap fixes the arrangement

			swaps++
			pairs.swap(i, pairs[i].idx)
			i--
		}
	}

	return swaps
}
