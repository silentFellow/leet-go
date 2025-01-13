package leetcode

import (
	"reflect"
	"testing"
)

func Test_commonChars(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example1",
			args: args{words: []string{"bella", "label", "roller"}},
			want: []string{"e", "l", "l"},
		},
		{
			name: "example2",
			args: args{words: []string{"cool", "lock", "cook"}},
			want: []string{"c", "o"},
		},
		{
			name: "no common characters",
			args: args{words: []string{"abc", "def", "ghi"}},
			want: []string{},
		},
		{
			name: "all characters common",
			args: args{words: []string{"aaa", "aaa", "aaa"}},
			want: []string{"a", "a", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMin(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name:  "all positive",
			args:  args{arr: []int{3, 1, 2}},
			want:  1,
			want1: true,
		},
		{
			name:  "contains zero",
			args:  args{arr: []int{3, 0, 2}},
			want:  3,
			want1: false,
		},
		{
			name:  "all same",
			args:  args{arr: []int{2, 2, 2}},
			want:  2,
			want1: true,
		},
		{
			name:  "single element",
			args:  args{arr: []int{5}},
			want:  5,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findMin(tt.args.arr)
			if got != tt.want {
				t.Errorf("findMin() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findMin() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
