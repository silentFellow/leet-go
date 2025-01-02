package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
  if head == nil || (head.Next == nil && n == 1) {
    return nil
  }

  if head.Next == nil && n != 1 {
    return head
  }

  ref := head
  for range n {
    ref = ref.Next
  }

  prev, deletable := head, head
  for ref != nil {
    prev = deletable
    deletable = deletable.Next
    ref = ref.Next
  }

  if deletable == head {
    return head.Next
  } else {
    prev.Next = deletable.Next
  }

  return head
}
