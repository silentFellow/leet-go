package leetcode

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example 1",
			args: args{root: &Node{Val: 1, Children: []*Node{
				{Val: 3, Children: []*Node{
					{Val: 5, Children: nil},
					{Val: 6, Children: nil},
				}},
				{Val: 2, Children: nil},
				{Val: 4, Children: nil},
			}}},
			want: [][]int{{1}, {3, 2, 4}, {5, 6}},
		},
		{
			name: "example 2",
			args: args{root: &Node{Val: 1, Children: []*Node{
				{Val: 2, Children: nil},
				{Val: 3, Children: []*Node{
					{Val: 6, Children: nil},
					{Val: 7, Children: []*Node{
						{Val: 11, Children: []*Node{
							{Val: 14, Children: nil},
						}},
					}},
					{Val: 8, Children: nil},
				}},
				{Val: 4, Children: nil},
				{Val: 5, Children: []*Node{
					{Val: 12, Children: nil},
					{Val: 13, Children: nil},
				}},
			}}},
			want: [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 12, 13}, {11}, {14}},
		},
		{
			name: "empty tree",
			args: args{root: nil},
			want: [][]int{},
		},
		{
			name: "single node",
			args: args{root: &Node{Val: 1, Children: nil}},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
