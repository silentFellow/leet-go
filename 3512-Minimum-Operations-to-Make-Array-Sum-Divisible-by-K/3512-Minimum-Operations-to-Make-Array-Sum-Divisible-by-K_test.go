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
			name: "sum already divisible",
			args: args{nums: []int{2, 4, 6}, k: 2},
			want: 0,
		},
		{
			name: "multiple operations needed",
			args: args{nums: []int{1, 1, 1}, k: 2},
			want: 1,
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
