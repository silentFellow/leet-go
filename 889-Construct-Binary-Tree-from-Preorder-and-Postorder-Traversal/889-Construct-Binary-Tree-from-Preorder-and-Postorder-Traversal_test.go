package leetcode

import (
	"reflect"
	"testing"
)

func Test_constructFromPrePost(t *testing.T) {
	type args struct {
		preorder  []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "example1",
			args: args{
				preorder:  []int{1, 2, 4, 5, 3, 6, 7},
				postorder: []int{4, 5, 2, 6, 7, 3, 1},
			},
			want: &TreeNode{
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
			},
		},
		{
			name: "example2",
			args: args{
				preorder:  []int{1},
				postorder: []int{1},
			},
			want: &TreeNode{
				Val: 1,
			},
		},
		{
			name: "example3",
			args: args{
				preorder:  []int{1, 2, 3},
				postorder: []int{3, 2, 1},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructFromPrePost(tt.args.preorder, tt.args.postorder); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("constructFromPrePost() = %v, want %v", got, tt.want)
			}
		})
	}
}
