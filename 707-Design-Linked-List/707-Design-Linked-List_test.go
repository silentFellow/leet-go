package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyLinkedList
	}{
		{
			name: "TestConstructor",
			want: MyLinkedList{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyLinkedList_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
		want int
	}{
		{
			name: "Get from empty list",
			this: &MyLinkedList{},
			args: args{index: 0},
			want: -1,
		},
		{
			name: "Get valid index",
			this: func() *MyLinkedList {
				ll := &MyLinkedList{}
				ll.AddAtHead(1)
				ll.AddAtTail(2)
				return ll
			}(),
			args: args{index: 1},
			want: 2,
		},
		{
			name: "Get invalid index",
			this: func() *MyLinkedList {
				ll := &MyLinkedList{}
				ll.AddAtHead(1)
				return ll
			}(),
			args: args{index: 2},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Get(tt.args.index); got != tt.want {
				t.Errorf("MyLinkedList.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyLinkedList_AddAtHead(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
		want []int
	}{
		{
			name: "Add at head",
			this: &MyLinkedList{},
			args: args{val: 1},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddAtHead(tt.args.val)
			if got := tt.this.Get(0); got != tt.want[0] {
				t.Errorf("MyLinkedList.AddAtHead() = %v, want %v", got, tt.want[0])
			}
		})
	}
}

func TestMyLinkedList_AddAtTail(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
		want []int
	}{
		{
			name: "Add at tail",
			this: func() *MyLinkedList {
				ll := &MyLinkedList{}
				ll.AddAtHead(1)
				return ll
			}(),
			args: args{val: 2},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddAtTail(tt.args.val)
			if got := tt.this.Get(1); got != tt.want[1] {
				t.Errorf("MyLinkedList.AddAtTail() = %v, want %v", got, tt.want[1])
			}
		})
	}
}

func TestMyLinkedList_AddAtIndex(t *testing.T) {
	type args struct {
		index int
		val   int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
		want []int
	}{
		{
			name: "Add at index",
			this: func() *MyLinkedList {
				ll := &MyLinkedList{}
				ll.AddAtHead(1)
				ll.AddAtTail(3)
				return ll
			}(),
			args: args{index: 1, val: 2},
			want: []int{1, 2, 3},
		},
		{
			name: "Add at invalid index",
			this: func() *MyLinkedList {
				ll := &MyLinkedList{}
				ll.AddAtHead(1)
				return ll
			}(),
			args: args{index: 2, val: 2},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddAtIndex(tt.args.index, tt.args.val)
			for i, v := range tt.want {
				if got := tt.this.Get(i); got != v {
					t.Errorf("MyLinkedList.AddAtIndex() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
		want []int
	}{
		{
			name: "Delete at index",
			this: func() *MyLinkedList {
				ll := &MyLinkedList{}
				ll.AddAtHead(1)
				ll.AddAtTail(2)
				ll.AddAtTail(3)
				return ll
			}(),
			args: args{index: 1},
			want: []int{1, 3},
		},
		{
			name: "Delete at invalid index",
			this: func() *MyLinkedList {
				ll := &MyLinkedList{}
				ll.AddAtHead(1)
				return ll
			}(),
			args: args{index: 1},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.DeleteAtIndex(tt.args.index)
			for i, v := range tt.want {
				if got := tt.this.Get(i); got != v {
					t.Errorf("MyLinkedList.DeleteAtIndex() = %v, want %v", got, v)
				}
			}
		})
	}
}
