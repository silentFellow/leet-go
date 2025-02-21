package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want FindElements
	}{
		{
			name: "single node",
			args: args{root: &TreeNode{Val: -1}},
			want: Constructor(&TreeNode{Val: 0}),
		},
		{
			name: "two levels",
			args: args{
				root: &TreeNode{Val: -1, Left: &TreeNode{Val: -1}, Right: &TreeNode{Val: -1}},
			},
			want: Constructor(&TreeNode{Val: 0, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.root); !reflect.DeepEqual(got.valMap, tt.want.valMap) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindElements_Find(t *testing.T) {
	type fields struct {
		valMap map[int]struct{}
	}
	type args struct {
		target int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "find existing element",
			fields: fields{valMap: map[int]struct{}{
				0: {}, 1: {}, 2: {},
			}},
			args: args{target: 1},
			want: true,
		},
		{
			name: "find non-existing element",
			fields: fields{valMap: map[int]struct{}{
				0: {}, 1: {}, 2: {},
			}},
			args: args{target: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &FindElements{
				valMap: tt.fields.valMap,
			}
			if got := this.Find(tt.args.target); got != tt.want {
				t.Errorf("FindElements.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_constructTree(t *testing.T) {
	type args struct {
		root   *TreeNode
		val    int
		valMap map[int]struct{}
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "single node",
			args: args{root: &TreeNode{Val: -1}, val: 0, valMap: map[int]struct{}{}},
			want: &TreeNode{Val: 0},
		},
		{
			name: "two levels",
			args: args{
				root:   &TreeNode{Val: -1, Left: &TreeNode{Val: -1}, Right: &TreeNode{Val: -1}},
				val:    0,
				valMap: map[int]struct{}{},
			},
			want: &TreeNode{Val: 0, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructTree(tt.args.root, tt.args.val, tt.args.valMap); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("constructTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
