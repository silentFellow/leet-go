package leetcode

import (
	"reflect"
	"testing"
)

func Test_getFinalState(t *testing.T) {
	type args struct {
		nums       []int
		k          int
		multiplier int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{nums: []int{2, 1, 3, 5, 6}, k: 5, multiplier: 2},
			want: []int{8, 4, 6, 5, 6},
		},
		{
			name: "example2",
			args: args{nums: []int{1, 2}, k: 3, multiplier: 4},
			want: []int{16, 8},
		},
		{
			name: "single_element",
			args: args{nums: []int{1}, k: 1, multiplier: 10},
			want: []int{10},
		},
		{
			name: "all_elements_same",
			args: args{nums: []int{2, 2, 2}, k: 2, multiplier: 3},
			want: []int{6, 6, 2},
		},
		{
			name: "no_operations",
			args: args{nums: []int{3, 1, 4, 1, 5}, k: 0, multiplier: 2},
			want: []int{3, 1, 4, 1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFinalState(tt.args.nums, tt.args.k, tt.args.multiplier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFinalState() = %v, want %v", got, tt.want)
			}
		})
	}
}
