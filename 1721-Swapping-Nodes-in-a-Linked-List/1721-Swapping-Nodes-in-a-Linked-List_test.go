package leetcode

import (
	"reflect"
	"testing"
)

func Test_swapNodes(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				k:    2,
			},
			want: &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, nil}}}}},
		},
		{
			name: "example 2",
			args: args{
				head: &ListNode{
					7,
					&ListNode{
						9,
						&ListNode{
							6,
							&ListNode{
								6,
								&ListNode{
									7,
									&ListNode{
										8,
										&ListNode{3, &ListNode{0, &ListNode{9, &ListNode{5, nil}}}},
									},
								},
							},
						},
					},
				},
				k: 5,
			},
			want: &ListNode{
				7,
				&ListNode{
					9,
					&ListNode{
						6,
						&ListNode{
							6,
							&ListNode{
								8,
								&ListNode{
									7,
									&ListNode{3, &ListNode{0, &ListNode{9, &ListNode{5, nil}}}},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "single node",
			args: args{
				head: &ListNode{1, nil},
				k:    1,
			},
			want: &ListNode{1, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapNodes(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
