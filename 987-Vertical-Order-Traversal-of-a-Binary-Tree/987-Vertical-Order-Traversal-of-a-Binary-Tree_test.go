package leetcode

import (
	"reflect"
	"testing"
)

func Test_verticalTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example1",
			args: args{root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			}},
			want: [][]int{{9}, {3, 15}, {20}, {7}},
		},
		{
			name: "example2",
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
			want: [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
		},
		{
			name: "example3",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			}},
			want: [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verticalTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
