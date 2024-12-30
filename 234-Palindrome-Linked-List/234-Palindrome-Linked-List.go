package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
  // identify the middle element
  slow, fast := head, head

  for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
  }

  // reverse the second half
  var prev, next *ListNode
  cur := slow

  for cur != nil {
    next = cur.Next
    cur.Next = prev
    prev = cur
    cur = next
  }

  // loop first and second
  // regular palindrome check
  first, second := head, prev
  for first != nil && second != nil {
    if first.Val != second.Val {
      return false
    }

    first = first.Next
    second = second.Next
  }

  return true
}
