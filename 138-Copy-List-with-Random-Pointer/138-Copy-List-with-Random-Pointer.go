package leetcode

// NOTE: try to do without map

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head
	for cur != nil {
		cur, cur.Next = cur.Next, &Node{Val: cur.Val, Next: cur.Next}
	}

	old := head
	for old != nil {
		copyNode := old.Next
		copyNode.Random = nil
		if old.Random != nil {
			copyNode.Random = old.Random.Next
		}
		old = copyNode.Next
	}

	newHead := head.Next
	old = head
	for old != nil {
		copyNode := old.Next
		oldNext := copyNode.Next
		copyNode.Next = nil
		if oldNext != nil {
			copyNode.Next = oldNext.Next
		}
		old.Next = oldNext
		old = old.Next
	}

	return newHead
}

// with map
// func copyRandomList(head *Node) *Node {
// 	curToNew := make(map[*Node]*Node)
//
// 	dummy := &Node{Val: 0}
// 	dummyTail := dummy
//
// 	th := head
// 	for th != nil {
// 		neo := &Node{Val: th.Val}
// 		dummyTail.Next = neo
// 		dummyTail = neo
// 		curToNew[th] = neo
// 		th = th.Next
// 	}
//
// 	th, nh := head, dummy.Next
// 	for th != nil && nh != nil {
// 		nh.Random = curToNew[th.Random]
// 		th = th.Next
// 		nh = nh.Next
// 	}
//
// 	return dummy.Next
// }
