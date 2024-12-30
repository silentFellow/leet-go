package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// return recursiveReverseList(nil, head)

  var prev, next *ListNode
	cur := head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

// for learning useing recursive method
func recursiveReverseList(prev *ListNode, cur *ListNode) *ListNode {
  if(cur == nil) {
    return prev
  }

  next := cur.Next
  cur.Next = prev

  return recursiveReverseList(cur, next)
}
