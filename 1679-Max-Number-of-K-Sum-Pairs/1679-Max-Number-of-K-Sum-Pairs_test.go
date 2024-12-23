package leetcode

import "testing"

func Test_maxOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    5,
			},
			want: 2,
		},
		{
			name: "example2",
			args: args{
				nums: []int{3, 1, 3, 4, 3},
				k:    6,
			},
			want: 1,
		},
		{
			name: "no pairs",
			args: args{
				nums: []int{1, 1, 1, 1},
				k:    3,
			},
			want: 0,
		},
		{
			name: "multiple pairs",
			args: args{
				nums: []int{2, 2, 2, 2},
				k:    4,
			},
			want: 2,
		},
		{
			name: "single element",
			args: args{
				nums: []int{1},
				k:    2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxOperations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
