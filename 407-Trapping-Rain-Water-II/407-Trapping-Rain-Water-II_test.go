package leetcode

import (
	"reflect"
	"testing"
)

func TestMinHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		h    MinHeap
		want int
	}{
		{name: "empty heap", h: MinHeap{}, want: 0},
		{name: "one element", h: MinHeap{{Height: 1, Row: 0, Col: 0}}, want: 1},
		{
			name: "multiple elements",
			h:    MinHeap{{Height: 1, Row: 0, Col: 0}, {Height: 2, Row: 1, Col: 1}},
			want: 2,
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
			name: "less",
			h:    MinHeap{{Height: 1, Row: 0, Col: 0}, {Height: 2, Row: 1, Col: 1}},
			args: struct{ i, j int }{i: 0, j: 1},
			want: true,
		},
		{
			name: "not less",
			h:    MinHeap{{Height: 2, Row: 0, Col: 0}, {Height: 1, Row: 1, Col: 1}},
			args: struct{ i, j int }{i: 0, j: 1},
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
			name: "swap",
			h:    MinHeap{{Height: 1, Row: 0, Col: 0}, {Height: 2, Row: 1, Col: 1}},
			args: struct{ i, j int }{i: 0, j: 1},
			want: MinHeap{{Height: 2, Row: 1, Col: 1}, {Height: 1, Row: 0, Col: 0}},
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
			name: "push",
			h:    &MinHeap{},
			args: struct{ x interface{} }{x: Cell{Height: 1, Row: 0, Col: 0}},
			want: MinHeap{{Height: 1, Row: 0, Col: 0}},
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
			name: "pop",
			h:    &MinHeap{{Height: 1, Row: 0, Col: 0}},
			want: Cell{Height: 1, Row: 0, Col: 0},
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

func Test_trapRainWater(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			grid [][]int
		}
		want int
	}{
		{
			name: "example 1",
			args: struct{ grid [][]int }{
				grid: [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}},
			},
			want: 4,
		},
		{
			name: "example 2",
			args: struct{ grid [][]int }{
				grid: [][]int{
					{3, 3, 3, 3, 3},
					{3, 2, 2, 2, 3},
					{3, 2, 1, 2, 3},
					{3, 2, 2, 2, 3},
					{3, 3, 3, 3, 3},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapRainWater(tt.args.grid); got != tt.want {
				t.Errorf("trapRainWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
