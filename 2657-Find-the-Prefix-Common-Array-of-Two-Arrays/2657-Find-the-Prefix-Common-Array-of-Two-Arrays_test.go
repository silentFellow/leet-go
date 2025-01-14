package leetcode

import (
	"reflect"
	"testing"
)

func Test_findThePrefixCommonArray(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{
			name: "test_case_1",
			args: args{
				A: []int{1, 3, 2, 4},
				B: []int{3, 1, 2, 4},
			},
			want: []int{0, 2, 3, 4},
		},
		{
			name: "test_case_2",
			args: args{
				A: []int{2, 3, 1},
				B: []int{3, 1, 2},
			},
			want: []int{0, 1, 3},
		},
		{
			name: "test_case_3",
			args: args{
				A: []int{1, 2, 3},
				B: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test_case_5",
			args: args{
				A: []int{1},
				B: []int{1},
			},
			want: []int{1},
		},
	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findThePrefixCommonArray(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findThePrefixCommonArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
