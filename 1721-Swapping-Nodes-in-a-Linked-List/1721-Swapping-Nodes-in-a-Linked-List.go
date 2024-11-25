package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
  slow, fast := head, head

  for range k-1 {
    fast = fast.Next
  }

  temp := fast

  for fast.Next != nil {
    fast = fast.Next
    slow = slow.Next
  }

  slow.Val, temp.Val = temp.Val, slow.Val

	return head
}
