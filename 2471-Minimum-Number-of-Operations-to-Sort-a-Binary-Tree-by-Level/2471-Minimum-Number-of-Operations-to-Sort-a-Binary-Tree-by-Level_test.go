package leetcode

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 9,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 10,
						},
					},
				},
			}},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			}},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
				},
			}},
			want: 0,
		},
		{
			name: "Single Node",
			args: args{root: &TreeNode{
				Val: 1,
			}},
			want: 0,
		},
		{
			name: "Two Levels",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 2,
				},
			}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.root); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
