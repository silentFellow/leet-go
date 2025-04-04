package leetcode

import (
	"reflect"
	"testing"
)

func Test_lcaDeepestLeaves(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{root: &TreeNode{Val: 1}},
			want: &TreeNode{Val: 1},
		},
		{
			name: "Example 2",
			args: args{root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 3},
			}},
			want: &TreeNode{Val: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcaDeepestLeaves(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lcaDeepestLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
