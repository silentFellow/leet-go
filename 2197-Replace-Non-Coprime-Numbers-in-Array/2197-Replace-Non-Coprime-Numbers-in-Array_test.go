package leetcode

import (
	"reflect"
	"testing"
)

func Test_replaceNonCoprimes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{6, 4, 3, 2, 7, 6, 2}},
			want: []int{12, 7, 6},
		},
		{
			name: "Example 2",
			args: args{nums: []int{2, 2, 1, 1, 3, 3, 3}},
			want: []int{2, 1, 1, 3},
		},
		{
			name: "Single element",
			args: args{nums: []int{5}},
			want: []int{5},
		},
		{
			name: "All coprime",
			args: args{nums: []int{2, 3, 5, 7}},
			want: []int{2, 3, 5, 7},
		},
		{
			name: "All non-coprime",
			args: args{nums: []int{4, 8, 16}},
			want: []int{16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceNonCoprimes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceNonCoprimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
