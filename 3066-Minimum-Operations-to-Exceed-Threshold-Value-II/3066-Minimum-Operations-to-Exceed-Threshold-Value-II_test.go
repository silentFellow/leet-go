package leetcode

import "testing"

func Test_minOperations(t *testing.T) {
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
			args: args{nums: []int{2, 11, 10, 1, 3}, k: 10},
			want: 2,
		},
		{
			name: "example2",
			args: args{nums: []int{1, 1, 2, 4, 9}, k: 20},
			want: 4,
		},
		{
			name: "all elements already greater than k",
			args: args{nums: []int{15, 20, 25}, k: 10},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
