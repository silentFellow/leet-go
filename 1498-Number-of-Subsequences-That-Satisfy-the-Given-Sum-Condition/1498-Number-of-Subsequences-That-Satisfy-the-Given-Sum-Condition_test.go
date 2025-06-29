package leetcode

import "testing"

func Test_numSubseq(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{nums: []int{3, 5, 6, 7}, target: 9},
			want: 4,
		},
		{
			name: "example2",
			args: args{nums: []int{3, 3, 6, 8}, target: 10},
			want: 6,
		},
		{
			name: "example3",
			args: args{nums: []int{2, 3, 3, 4, 6, 7}, target: 12},
			want: 61,
		},
		{
			name: "single element valid",
			args: args{nums: []int{1}, target: 2},
			want: 1,
		},
		{
			name: "single element invalid",
			args: args{nums: []int{5}, target: 4},
			want: 0,
		},
		{
			name: "all pairs valid",
			args: args{nums: []int{1, 1, 1}, target: 3},
			want: 7,
		},
		{
			name: "no valid subsequence",
			args: args{nums: []int{10, 11, 12}, target: 5},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubseq(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("numSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
