package leetcode

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "symmetric tree",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			}},
			want: true,
		},
		{
			name: "asymmetric tree",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			}},
			want: false,
		},
		{
			name: "single node tree",
			args: args{root: &TreeNode{
				Val: 1,
			}},
			want: true,
		},
		{
			name: "empty tree",
			args: args{root: nil},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
