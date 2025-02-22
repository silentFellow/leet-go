package leetcode

import (
	"reflect"
	"testing"
)

func Test_recoverFromPreorder(t *testing.T) {
	type args struct {
		traversal string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "example1",
			args: args{traversal: "1-2--3--4-5--6--7"},
			want: &TreeNode{
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
					Val: 5,
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
			args: args{traversal: "1-2--3---4-5--6---7"},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 7,
						},
					},
				},
			},
		},
		{
			name: "example3",
			args: args{traversal: "1-401--349---90--88"},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 401,
					Left: &TreeNode{
						Val: 349,
						Left: &TreeNode{
							Val: 90,
						},
					},
					Right: &TreeNode{
						Val: 88,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recoverFromPreorder(tt.args.traversal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recoverFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
