package leetcode

import (
	"reflect"
	"testing"
)

func TestProductOfNumbers_Add(t *testing.T) {
	tests := []struct {
		name   string
		fields []int
		args   int
		want   []int
	}{
		{
			name:   "Add 3 to empty list",
			fields: []int{},
			args:   3,
			want:   []int{3},
		},
		{
			name:   "Add 0 to list",
			fields: []int{3},
			args:   0,
			want:   []int{3, 0},
		},
		{
			name:   "Add 2 to list",
			fields: []int{3, 0},
			args:   2,
			want:   []int{3, 0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &ProductOfNumbers{
				nums: tt.fields,
			}
			this.Add(tt.args)
			if !reflect.DeepEqual(this.nums, tt.want) {
				t.Errorf("ProductOfNumbers.Add() = %v, want %v", this.nums, tt.want)
			}
		})
	}
}

func TestProductOfNumbers_GetProduct(t *testing.T) {
	tests := []struct {
		name   string
		fields []int
		args   int
		want   int
	}{
		{
			name:   "Get product of last 2 numbers",
			fields: []int{3, 0, 2, 5, 4},
			args:   2,
			want:   20,
		},
		{
			name:   "Get product of last 3 numbers",
			fields: []int{3, 0, 2, 5, 4},
			args:   3,
			want:   40,
		},
		{
			name:   "Get product of last 4 numbers",
			fields: []int{3, 0, 2, 5, 4},
			args:   4,
			want:   0,
		},
		{
			name:   "Get product of last 2 numbers after adding 8",
			fields: []int{3, 0, 2, 5, 4, 8},
			args:   2,
			want:   32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &ProductOfNumbers{
				nums: tt.fields,
			}
			if got := this.GetProduct(tt.args); got != tt.want {
				t.Errorf("ProductOfNumbers.GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
