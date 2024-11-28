package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
  // find the first non val node
  for head != nil && head.Val == val {
    head = head.Next
  }

	cur, prev := head, head

	for cur != nil {
		for cur != nil && cur.Val != val {
			prev = cur
			cur = cur.Next
		}

		if cur != nil {
			prev.Next = cur.Next
      cur = cur.Next
		}
	}

	return head
}
