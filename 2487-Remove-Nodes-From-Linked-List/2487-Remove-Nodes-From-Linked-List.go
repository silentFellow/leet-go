package leetcode

// TODO: new to learn and improvise it

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
  stack := []*ListNode{}

  for head != nil {
    for len(stack) > 0 && head.Val > stack[len(stack)-1].Val {
      stack = stack[:len(stack)-1]
    }

    // to avoid nil pointer error if stack is empty
    if len(stack) > 0 {
      stack[len(stack) - 1].Next = head // connect the previous and current node in stack
    }
    stack = append(stack, head)
    head = head.Next
  }

  return stack[0]
}
