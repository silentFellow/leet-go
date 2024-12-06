package leetcode

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "target in matrix",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "target not in matrix",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 13,
			},
			want: false,
		},
		{
			name: "target is the first element",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 1,
			},
			want: true,
		},
		{
			name: "target is the last element",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 60,
			},
			want: true,
		},
		{
			name: "single row matrix",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
				},
				target: 5,
			},
			want: true,
		},
		{
			name: "single column matrix",
			args: args{
				matrix: [][]int{
					{1},
					{3},
					{5},
					{7},
				},
				target: 5,
			},
			want: true,
		},
		{
			name: "negative numbers in matrix",
			args: args{
				matrix: [][]int{
					{-10, -5, 0, 5},
					{10, 15, 20, 25},
					{30, 35, 40, 45},
				},
				target: -5,
			},
			want: true,
		},
		{
			name: "target not in matrix with negative numbers",
			args: args{
				matrix: [][]int{
					{-10, -5, 0, 5},
					{10, 15, 20, 25},
					{30, 35, 40, 45},
				},
				target: -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
