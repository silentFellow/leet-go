package leetcode

import (
	"reflect"
	"testing"
)

func Test_prefixesDivBy5(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "Example 1",
			args: args{nums: []int{0, 1, 1}},
			want: []bool{true, false, false},
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 1, 1}},
			want: []bool{false, false, false},
		},
		{
			name: "Single zero",
			args: args{nums: []int{0}},
			want: []bool{true},
		},
		{
			name: "Single one",
			args: args{nums: []int{1}},
			want: []bool{false},
		},
		{
			name: "Multiple divisible",
			args: args{nums: []int{1, 0, 1}},
			want: []bool{false, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixesDivBy5(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixesDivBy5() = %v, want %v", got, tt.want)
			}
		})
	}
}
