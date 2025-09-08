package leetcode

import (
	"reflect"
	"testing"
)

func Test_getNoZeroIntegers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"n=2", args{2}, []int{1, 1}},
		{"n=11", args{11}, []int{2, 9}},
		{"n=100", args{100}, []int{1, 99}},
		{"n=101", args{101}, []int{2, 99}},
		{"n=1000", args{1000}, []int{1, 999}},
		{"n=19", args{19}, []int{1, 18}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNoZeroIntegers(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNoZeroIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
