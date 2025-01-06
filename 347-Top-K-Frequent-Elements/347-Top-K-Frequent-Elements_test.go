package leetcode

import (
	"reflect"
	"testing"
)

func Test_maxHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		m    maxHeap
		want int
	}{
		{name: "empty heap", m: maxHeap{}, want: 0},
		{name: "one element", m: maxHeap{{val: 1, freq: 1}}, want: 1},
		{name: "multiple elements", m: maxHeap{{val: 1, freq: 3}, {val: 2, freq: 2}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Len(); got != tt.want {
				t.Errorf("maxHeap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxHeap_Swap(t *testing.T) {
	tests := []struct {
		name string
		m    maxHeap
		args struct{ i, j int }
		want maxHeap
	}{
		{name: "swap elements", m: maxHeap{{val: 1, freq: 3}, {val: 2, freq: 2}}, args: struct{ i, j int }{0, 1}, want: maxHeap{{val: 2, freq: 2}, {val: 1, freq: 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Swap(tt.args.i, tt.args.j)
			if !reflect.DeepEqual(tt.m, tt.want) {
				t.Errorf("maxHeap.Swap() = %v, want %v", tt.m, tt.want)
			}
		})
	}
}

func Test_maxHeap_Less(t *testing.T) {
	tests := []struct {
		name string
		m    maxHeap
		args struct{ i, j int }
		want bool
	}{
		{name: "compare elements", m: maxHeap{{val: 1, freq: 3}, {val: 2, freq: 2}}, args: struct{ i, j int }{0, 1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("maxHeap.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxHeap_Push(t *testing.T) {
	tests := []struct {
		name string
		m    *maxHeap
		args pairs
		want maxHeap
	}{
		{name: "push element", m: &maxHeap{}, args: pairs{val: 1, freq: 3}, want: maxHeap{{val: 1, freq: 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Push(tt.args)
			if !reflect.DeepEqual(*tt.m, tt.want) {
				t.Errorf("maxHeap.Push() = %v, want %v", *tt.m, tt.want)
			}
		})
	}
}

func Test_maxHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		m    *maxHeap
		want pairs
	}{
		{name: "pop element", m: &maxHeap{{val: 1, freq: 3}}, want: pairs{val: 1, freq: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			nums []int
			k    int
		}
		want []int
	}{
		{name: "example 1", args: struct{ nums []int; k int }{nums: []int{1, 1, 1, 2, 2, 3}, k: 2}, want: []int{1, 2}},
		{name: "example 2", args: struct{ nums []int; k int }{nums: []int{1}, k: 1}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
