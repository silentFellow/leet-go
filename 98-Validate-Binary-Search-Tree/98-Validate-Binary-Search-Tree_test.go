package leetcode

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				root: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
				},
			},
			want: false,
		},
		{
			name: "Single Node",
			args: args{root: &TreeNode{Val: 1}},
			want: true,
		},
		{
			name: "Invalid BST with equal values",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}},
			want: false,
		},
		{
			name: "Valid BST with negative values",
			args: args{root: &TreeNode{Val: 0, Left: &TreeNode{Val: -1}, Right: &TreeNode{Val: 1}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
