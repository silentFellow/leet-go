package leetcode

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "balanced tree",
			args: args{
				root: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
				},
			},
			want: true,
		},
		{
			name: "unbalanced tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 2},
				},
			},
			want: false,
		},
		{
			name: "empty tree",
			args: args{root: nil},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "depth of tree with one node",
			args: args{root: &TreeNode{Val: 1}},
			want: 1,
		},
		{
			name: "depth of tree with multiple nodes",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
					Right: &TreeNode{Val: 5},
				},
			},
			want: 3,
		},
		{
			name: "depth of empty tree",
			args: args{root: nil},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDepth(tt.args.root); got != tt.want {
				t.Errorf("getDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive number",
			args: args{v: 5},
			want: 5,
		},
		{
			name: "negative number",
			args: args{v: -5},
			want: 5,
		},
		{
			name: "zero",
			args: args{v: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.v); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
