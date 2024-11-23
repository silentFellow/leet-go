package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
  // find the middle element
  steps := 0
  slow, fast := head, head
  for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
    steps += 1
  }

  // reverse from middle element
  var prev *ListNode
  for slow != nil {
    next := slow.Next
    slow.Next = prev
    prev = slow
    slow = next
  }

  maxSum := 0
  for range steps {
    sum := head.Val + prev.Val
    if maxSum < sum {
      maxSum = sum
    }

    head = head.Next
    prev = prev.Next
  }

  return maxSum
}
