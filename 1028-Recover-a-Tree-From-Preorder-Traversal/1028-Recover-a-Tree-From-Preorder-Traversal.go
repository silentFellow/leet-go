package leetcode

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(traversal string) *TreeNode {
	dash := 0
	stack := []*TreeNode{}
	i := 0

	for i < len(traversal) {
    // get no. of dashes
		for i < len(traversal) && traversal[i] == '-' {
			dash++
			i++
		}

    // get the int value
		start := i
		for i < len(traversal) && traversal[i] >= '0' && traversal[i] <= '9' {
			i++
		}
		v, _ := strconv.Atoi(traversal[start:i])

    // check the proper level
		for len(stack) != dash {
			stack = stack[:len(stack)-1]
		}

		newNode := &TreeNode{Val: v}
		if len(stack) != 0 && stack[len(stack)-1].Left == nil {
			stack[len(stack)-1].Left = newNode
		} else if len(stack) != 0 && stack[len(stack)-1].Right == nil {
			stack[len(stack)-1].Right = newNode
		}
		stack = append(stack, newNode)

		dash = 0
	}

	return stack[0]
}
