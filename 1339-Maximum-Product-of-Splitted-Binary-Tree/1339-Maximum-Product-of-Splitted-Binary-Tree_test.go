package leetcode

import "testing"

func buildTreeFromSlice(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v != nil {
			nodes[i] = &TreeNode{Val: v.(int)}
		}
	}
	for i, node := range nodes {
		if node == nil {
			continue
		}
		leftIdx := 2*i + 1
		rightIdx := 2*i + 2
		if leftIdx < len(vals) {
			node.Left = nodes[leftIdx]
		}
		if rightIdx < len(vals) {
			node.Right = nodes[rightIdx]
		}
	}
	return nodes[0]
}

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"Example1", buildTreeFromSlice([]any{1, 2, 3, 4, 5, 6}), 110},
		{"Example2", buildTreeFromSlice([]any{1, nil, 2, 3, 4, nil, nil, 5, 6}), 2},
		{"TwoNodes", buildTreeFromSlice([]any{2, 3}), 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.root); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
