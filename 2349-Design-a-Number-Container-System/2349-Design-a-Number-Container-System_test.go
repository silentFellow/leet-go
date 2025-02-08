package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want NumberContainers
	}{
		{
			name: "Test Constructor",
			want: NumberContainers{
				indexMap: map[int]int{},
				hmap:     map[int]*MinHeap{},
			},
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

func TestNumberContainers_Change(t *testing.T) {
	tests := []struct {
		name   string
		fields NumberContainers
		args   struct {
			index  int
			number int
		}
		want map[int]int
	}{
		{
			name: "Test Change 1",
			fields: NumberContainers{
				indexMap: map[int]int{},
				hmap:     map[int]*MinHeap{},
			},
			args: struct {
				index  int
				number int
			}{index: 1, number: 10},
			want: map[int]int{1: 10},
		},
		{
			name: "Test Change 2",
			fields: NumberContainers{
				indexMap: map[int]int{1: 10},
				hmap:     map[int]*MinHeap{10: {1}},
			},
			args: struct {
				index  int
				number int
			}{index: 1, number: 20},
			want: map[int]int{1: 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &NumberContainers{
				indexMap: tt.fields.indexMap,
				hmap:     tt.fields.hmap,
			}
			this.Change(tt.args.index, tt.args.number)
			if !reflect.DeepEqual(this.indexMap, tt.want) {
				t.Errorf("NumberContainers.Change() = %v, want %v", this.indexMap, tt.want)
			}
		})
	}
}

func TestNumberContainers_Find(t *testing.T) {
	tests := []struct {
		name   string
		fields NumberContainers
		args   struct {
			number int
		}
		want int
	}{
		{
			name: "Test Find 1",
			fields: NumberContainers{
				indexMap: map[int]int{1: 10, 2: 10},
				hmap:     map[int]*MinHeap{10: {1, 2}},
			},
			args: struct {
				number int
			}{number: 10},
			want: 1,
		},
		{
			name: "Test Find 2",
			fields: NumberContainers{
				indexMap: map[int]int{1: 10, 2: 20},
				hmap:     map[int]*MinHeap{10: {1}, 20: {2}},
			},
			args: struct {
				number int
			}{number: 20},
			want: 2,
		},
		{
			name: "Test Find 3",
			fields: NumberContainers{
				indexMap: map[int]int{1: 10, 2: 20},
				hmap:     map[int]*MinHeap{10: {1}, 20: {2}},
			},
			args: struct {
				number int
			}{number: 30},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &NumberContainers{
				indexMap: tt.fields.indexMap,
				hmap:     tt.fields.hmap,
			}
			if got := this.Find(tt.args.number); got != tt.want {
				t.Errorf("NumberContainers.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		h    MinHeap
		want int
	}{
		{
			name: "Test Len 1",
			h:    MinHeap{1, 2, 3},
			want: 3,
		},
		{
			name: "Test Len 2",
			h:    MinHeap{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("MinHeap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Less(t *testing.T) {
	tests := []struct {
		name string
		h    MinHeap
		args struct {
			i int
			j int
		}
		want bool
	}{
		{
			name: "Test Less 1",
			h:    MinHeap{1, 2, 3},
			args: struct {
				i int
				j int
			}{i: 0, j: 1},
			want: true,
		},
		{
			name: "Test Less 2",
			h:    MinHeap{3, 2, 1},
			args: struct {
				i int
				j int
			}{i: 0, j: 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("MinHeap.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Swap(t *testing.T) {
	tests := []struct {
		name string
		h    MinHeap
		args struct {
			i int
			j int
		}
		want MinHeap
	}{
		{
			name: "Test Swap 1",
			h:    MinHeap{1, 2, 3},
			args: struct {
				i int
				j int
			}{i: 0, j: 2},
			want: MinHeap{3, 2, 1},
		},
		{
			name: "Test Swap 2",
			h:    MinHeap{3, 2, 1},
			args: struct {
				i int
				j int
			}{i: 1, j: 2},
			want: MinHeap{3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
			if !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("MinHeap.Swap() = %v, want %v", tt.h, tt.want)
			}
		})
	}
}

func TestMinHeap_Push(t *testing.T) {
	tests := []struct {
		name string
		h    *MinHeap
		args struct {
			x interface{}
		}
		want MinHeap
	}{
		{
			name: "Test Push 1",
			h:    &MinHeap{1, 2},
			args: struct {
				x interface{}
			}{x: 3},
			want: MinHeap{1, 2, 3},
		},
		{
			name: "Test Push 2",
			h:    &MinHeap{},
			args: struct {
				x interface{}
			}{x: 1},
			want: MinHeap{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.args.x)
			if !reflect.DeepEqual(*tt.h, tt.want) {
				t.Errorf("MinHeap.Push() = %v, want %v", *tt.h, tt.want)
			}
		})
	}
}

func TestMinHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    *MinHeap
		want interface{}
	}{
		{
			name: "Test Pop 1",
			h:    &MinHeap{1, 2, 3},
			want: 3,
		},
		{
			name: "Test Pop 2",
			h:    &MinHeap{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Remove(t *testing.T) {
	tests := []struct {
		name string
		h    *MinHeap
		args struct {
			index int
		}
		want MinHeap
	}{
		{
			name: "Test Remove 1",
			h:    &MinHeap{1, 2, 3},
			args: struct {
				index int
			}{index: 2},
			want: MinHeap{1, 3},
		},
		{
			name: "Test Remove 2",
			h:    &MinHeap{1},
			args: struct {
				index int
			}{index: 1},
			want: MinHeap{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Remove(tt.args.index)
			if !reflect.DeepEqual(*tt.h, tt.want) {
				t.Errorf("MinHeap.Remove() = %v, want %v", *tt.h, tt.want)
			}
		})
	}
}
