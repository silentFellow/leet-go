package leetcode

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "example 1",
			args: args{head: &Node{Val: 7, Next: &Node{Val: 13, Next: &Node{Val: 11, Next: &Node{Val: 10, Next: &Node{Val: 1}}}}}},
			want: &Node{Val: 7, Next: &Node{Val: 13, Next: &Node{Val: 11, Next: &Node{Val: 10, Next: &Node{Val: 1}}}}}, // Expected deep copy structure
		},
		{
			name: "example 2",
			args: args{head: &Node{Val: 1, Next: &Node{Val: 2}}},
			want: &Node{Val: 1, Next: &Node{Val: 2}}, // Expected deep copy structure
		},
		{
			name: "example 3",
			args: args{head: &Node{Val: 3, Next: &Node{Val: 3, Next: &Node{Val: 3}}}},
			want: &Node{Val: 3, Next: &Node{Val: 3, Next: &Node{Val: 3}}}, // Expected deep copy structure
		},
		{
			name: "empty list",
			args: args{head: nil},
			want: nil, // Expected deep copy structure
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}
