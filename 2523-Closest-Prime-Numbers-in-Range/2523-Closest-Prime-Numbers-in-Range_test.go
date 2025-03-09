package leetcode

import (
	"reflect"
	"testing"
)

func Test_closestPrimes(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test case 1",
			args: args{left: 10, right: 19},
			want: []int{11, 13},
		},
		{
			name: "Test case 2",
			args: args{left: 4, right: 6},
			want: []int{-1, -1},
		},
		{
			name: "Test case 3",
			args: args{left: 1, right: 10},
			want: []int{2, 3},
		},
		{
			name: "Test case 4",
			args: args{left: 14, right: 20},
			want: []int{17, 19},
		},
		{
			name: "Test case 5",
			args: args{left: 22, right: 29},
			want: []int{23, 29},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestPrimes(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
