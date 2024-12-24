package leetcode

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty tree",
			args: args{root: nil},
			want: 0,
		},
		{
			name: "single node",
			args: args{root: &TreeNode{Val: 1}},
			want: 0,
		},
		{
			name: "two nodes",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}},
			want: 1,
		},
		{
			name: "three nodes",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}},
			want: 2,
		},
		{
			name: "complex tree",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
