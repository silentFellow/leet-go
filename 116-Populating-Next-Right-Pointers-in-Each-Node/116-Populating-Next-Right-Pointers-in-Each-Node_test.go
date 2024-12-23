package leetcode

import (
	"fmt"
	"testing"
)

func isTreeEqual(node1, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		fmt.Printf("Mismatch: node1 = %+v, node2 = %+v\n", node1, node2)
		return false
	}
	if node1.Val != node2.Val {
		fmt.Printf("Value mismatch: node1.Val = %d, node2.Val = %d\n", node1.Val, node2.Val)
		return false
	}
	if (node1.Next == nil) != (node2.Next == nil) {
		fmt.Printf("Next mismatch: node1.Next = %+v, node2.Next = %+v\n", node1.Next, node2.Next)
		return false
	}
	if node1.Next != nil && node1.Next.Val != node2.Next.Val {
		fmt.Printf(
			"Next value mismatch: node1.Next.Val = %d, node2.Next.Val = %d\n",
			node1.Next.Val,
			node2.Next.Val,
		)
		return false
	}
	return isTreeEqual(node1.Left, node2.Left) &&
		isTreeEqual(node1.Right, node2.Right)
}

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Single Node",
			args: args{root: &Node{Val: 1}},
			want: &Node{Val: 1},
		},
		{
			name: "Perfect Binary Tree",
			args: args{
				root: &Node{
					Val: 1,
					Left: &Node{
						Val:   2,
						Left:  &Node{Val: 4},
						Right: &Node{Val: 5},
					},
					Right: &Node{
						Val:   3,
						Left:  &Node{Val: 6},
						Right: &Node{Val: 7},
					},
				},
			},
			want: &Node{
				Val: 1,
				Left: &Node{
					Val:   2,
					Next:  &Node{Val: 3},
					Left:  &Node{Val: 4, Next: &Node{Val: 5}},
					Right: &Node{Val: 5, Next: &Node{Val: 6}},
				},
				Right: &Node{
					Val:   3,
					Left:  &Node{Val: 6, Next: &Node{Val: 7}},
					Right: &Node{Val: 7},
				},
			},
		},
		{
			name: "Imbalanced Tree",
			args: args{
				root: &Node{
					Val:  1,
					Left: &Node{Val: 2, Right: &Node{Val: 5}},
				},
			},
			want: &Node{
				Val:  1,
				Left: &Node{Val: 2, Next: nil, Right: &Node{Val: 5}},
			},
		},
		{
			name: "Empty Tree",
			args: args{root: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := connect(tt.args.root)
			if !isTreeEqual(got, tt.want) {
				t.Errorf("connect() failed for test %s\n", tt.name)
				fmt.Println("Expected:")
				printTree(tt.want)
				fmt.Println("Got:")
				printTree(got)
			}
		})
	}
}

func printTree(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("Node %d, Next: %v\n", root.Val, getNextVal(root.Next))
	printTree(root.Left)
	printTree(root.Right)
}

func getNextVal(node *Node) interface{} {
	if node == nil {
		return nil
	}
	return node.Val
}
