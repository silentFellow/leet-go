package leetcode

import "testing"

func Test_minimumSize(t *testing.T) {
	type args struct {
		nums          []int
		maxOperations int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums:          []int{9},
				maxOperations: 2,
			},
			want: 3,
		},
		{
			name: "example2",
			args: args{
				nums:          []int{2, 4, 8, 2},
				maxOperations: 4,
			},
			want: 2,
		},
		{
			name: "example3",
			args: args{
				nums:          []int{7, 17},
				maxOperations: 2,
			},
			want: 7,
		},
		{
			name: "example4",
			args: args{
				nums:          []int{1, 1, 1, 1},
				maxOperations: 0,
			},
			want: 1,
		},
		{
			name: "example5",
			args: args{
				nums:          []int{10, 10, 10},
				maxOperations: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSize(tt.args.nums, tt.args.maxOperations); got != tt.want {
				t.Errorf("minimumSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
