package leetcode

type ListNode struct {
	Next *ListNode
	Val  int
}

type MyLinkedList struct {
	head   *ListNode
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index+1 > this.length {
		return -1
	}

	temp := this.head
	for range index {
		temp = temp.Next
	}

	return temp.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	defer func() {
		this.length++
	}()

	neo := &ListNode{
		Val: val,
	}

	if this.head == nil {
		this.head = neo
		return
	}

	neo.Next = this.head
	this.head = neo
}

func (this *MyLinkedList) AddAtTail(val int) {
	defer func() {
		this.length++
	}()

	neo := &ListNode{Val: val}
	if this.head == nil {
		this.head = neo
		return
	}

	temp := this.head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = neo
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.length {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	} else if index == this.length {
		this.AddAtTail(val)
		return
	}

  // cautious on using defer, if not inserted should not increase. if used in if statement 2 defer occurs
	defer func() {
		this.length++
	}()

	neo := &ListNode{Val: val}
	temp := this.head
	for range index - 1 {
		temp = temp.Next
	}
	neo.Next = temp.Next
	temp.Next = neo
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index+1 > this.length {
		return
	}

	defer func() {
		this.length--
	}()

	if index == 0 {
		this.head = this.head.Next
		return
	}

	temp := this.head
	for range index - 1 {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
