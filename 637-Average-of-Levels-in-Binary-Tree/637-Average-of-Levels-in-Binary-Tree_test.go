package leetcode

import (
	"reflect"
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "example 1",
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
			want: []float64{3.00000, 14.50000, 11.00000},
		},
		{
			name: "example 2",
			args: args{root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 20,
				},
			}},
			want: []float64{3.00000, 14.50000, 11.00000},
		},
		{
			name: "single node",
			args: args{root: &TreeNode{
				Val: 1,
			}},
			want: []float64{1.00000},
		},
		{
			name: "empty tree",
			args: args{root: nil},
			want: []float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
