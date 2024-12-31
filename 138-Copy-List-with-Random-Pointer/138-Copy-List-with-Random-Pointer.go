package leetcode

// NOTE: try to do without map

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	curToNew := make(map[*Node]*Node)

	dummy := &Node{Val: 0}
	dummyTail := dummy

	th := head
	for th != nil {
		neo := &Node{Val: th.Val}
		dummyTail.Next = neo
		dummyTail = neo
		curToNew[th] = neo
		th = th.Next
	}

	th, nh := head, dummy.Next
	for th != nil && nh != nil {
		nh.Random = curToNew[th.Random]
		th = th.Next
		nh = nh.Next
	}

	return dummy.Next
}
