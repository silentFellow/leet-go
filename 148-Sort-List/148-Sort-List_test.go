package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{head: &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		},
		{
			name: "example 2",
			args: args{
				head: &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}},
			},
			want: &ListNode{-1, &ListNode{0, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
		{
			name: "empty list",
			args: args{head: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		f *ListNode
		s *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "merge two sorted lists",
			args: args{
				f: &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
				s: &ListNode{2, &ListNode{4, &ListNode{6, nil}}},
			},
			want: &ListNode{
				1,
				&ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}},
			},
		},
		{
			name: "merge with one empty list",
			args: args{
				f: &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
				s: nil,
			},
			want: &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
		},
		{
			name: "merge two empty lists",
			args: args{
				f: nil,
				s: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.f, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
