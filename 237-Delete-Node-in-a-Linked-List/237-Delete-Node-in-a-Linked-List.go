package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
  prev := node
  for node.Next != nil {
    temp := node.Val
    node.Val = node.Next.Val
    node.Next.Val = temp

    prev = node
    node = node.Next
  }
  
  prev.Next = nil
}
