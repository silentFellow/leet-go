package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeInBetween(t *testing.T) {
	type args struct {
		list1 *ListNode
		a     int
		b     int
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				list1: &ListNode{10, &ListNode{1, &ListNode{13, &ListNode{6, &ListNode{9, &ListNode{5, nil}}}}}},
				a:     3,
				b:     4,
				list2: &ListNode{1000000, &ListNode{1000001, &ListNode{1000002, nil}}},
			},
			want: &ListNode{10, &ListNode{1, &ListNode{13, &ListNode{1000000, &ListNode{1000001, &ListNode{1000002, &ListNode{5, nil}}}}}}},
		},
		{
			name: "test2",
			args: args{
				list1: &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}},
				a:     2,
				b:     5,
				list2: &ListNode{1000000, &ListNode{1000001, &ListNode{1000002, &ListNode{1000003, &ListNode{1000004, nil}}}}},
			},
			want: &ListNode{0, &ListNode{1, &ListNode{1000000, &ListNode{1000001, &ListNode{1000002, &ListNode{1000003, &ListNode{1000004, &ListNode{6, nil}}}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeInBetween(tt.args.list1, tt.args.a, tt.args.b, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeInBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
