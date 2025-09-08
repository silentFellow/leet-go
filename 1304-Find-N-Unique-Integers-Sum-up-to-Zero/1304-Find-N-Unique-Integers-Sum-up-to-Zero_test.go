package leetcode

import (
	"reflect"
	"testing"
)

func Test_sumZero(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"n=1", args{1}, []int{0}},
		{"n=2", args{2}, []int{-1, 1}},
		{"n=3", args{3}, []int{-1, 0, 1}},
		{"n=4", args{4}, []int{-2, -1, 1, 2}},
		{"n=5", args{5}, []int{-2, -1, 0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumZero(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
