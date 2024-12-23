package leetcode

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			}},
			want: []int{1, 3, 4},
		},
		{
			name: "Example 2",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			}},
			want: []int{1, 3, 4},
		},
		{
			name: "Example 3",
			args: args{root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 3,
				},
			}},
			want: []int{1, 3},
		},
		{
			name: "Example 4",
			args: args{root: nil},
			want: []int{},
		},
		{
			name: "Single Node",
			args: args{root: &TreeNode{
				Val: 1,
			}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
