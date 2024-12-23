package leetcode

import "testing"

func Test_maxLevelSum(t *testing.T) {
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
					Val: 7,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: -8,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			}},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{root: &TreeNode{
				Val: 989,
				Right: &TreeNode{
					Val: 10250,
					Left: &TreeNode{
						Val: 98693,
					},
					Right: &TreeNode{
						Val: -89388,
						Right: &TreeNode{
							Val: -32127,
						},
					},
				},
			}},
			want: 2,
		},
		{
			name: "Single Node",
			args: args{root: &TreeNode{
				Val: 1,
			}},
			want: 1,
		},
		{
			name: "All Negative Values",
			args: args{root: &TreeNode{
				Val: -1,
				Left: &TreeNode{
					Val: -2,
					Left: &TreeNode{
						Val: -4,
					},
					Right: &TreeNode{
						Val: -5,
					},
				},
				Right: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -6,
					},
					Right: &TreeNode{
						Val: -7,
					},
				},
			}},
			want: 1,
		},
		{
			name: "Mixed Values",
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
					Right: &TreeNode{
						Val: 7,
					},
				},
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLevelSum(tt.args.root); got != tt.want {
				t.Errorf("maxLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
