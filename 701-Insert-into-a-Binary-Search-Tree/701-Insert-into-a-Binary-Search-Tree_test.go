package leetcode

import (
	"reflect"
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Insert into empty tree",
			args: args{
				root: nil,
				val:  5,
			},
			want: &TreeNode{
				Val: 5,
			},
		},
		{
			name: "Insert into left subtree",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 15,
					},
				},
				val: 3,
			},
			want: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 15,
				},
			},
		},
		{
			name: "Insert into right subtree",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 15,
					},
				},
				val: 20,
			},
			want: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 15,
					Right: &TreeNode{
						Val: 20,
					},
				},
			},
		},
		{
			name: "Insert into deeper level",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				val: 5,
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
