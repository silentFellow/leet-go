package leetcode

import "testing"

func Test_maxCount(t *testing.T) {
	type args struct {
		banned []int
		n      int
		maxSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				banned: []int{1, 6, 5},
				n:      5,
				maxSum: 6,
			},
			want: 2,
		},
		{
			name: "example2",
			args: args{
				banned: []int{1, 2, 3, 4, 5, 6, 7},
				n:      8,
				maxSum: 1,
			},
			want: 0,
		},
		{
			name: "example3",
			args: args{
				banned: []int{11},
				n:      7,
				maxSum: 50,
			},
			want: 7,
		},
		{
			name: "no banned numbers",
			args: args{
				banned: []int{},
				n:      10,
				maxSum: 15,
			},
			want: 5,
		},
		{
			name: "all numbers banned",
			args: args{
				banned: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				n:      10,
				maxSum: 15,
			},
			want: 0,
		},
		{
			name: "maxSum too low",
			args: args{
				banned: []int{2, 4, 6, 8, 10},
				n:      10,
				maxSum: 1,
			},
			want: 1,
		},
		{
			name: "maxSum just enough",
			args: args{
				banned: []int{2, 4, 6, 8, 10},
				n:      10,
				maxSum: 15,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCount(tt.args.banned, tt.args.n, tt.args.maxSum); got != tt.want {
				t.Errorf("maxCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
