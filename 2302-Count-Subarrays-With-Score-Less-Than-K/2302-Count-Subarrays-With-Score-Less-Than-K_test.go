package leetcode

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example1",
			args: args{
				nums: []int{2, 1, 4, 3, 5},
				k:    10,
			},
			want: 6,
		},
		{
			name: "example2",
			args: args{
				nums: []int{1, 1, 1},
				k:    5,
			},
			want: 5,
		},
		{
			name: "single_element",
			args: args{
				nums: []int{10},
				k:    15,
			},
			want: 1,
		},
		{
			name: "no_valid_subarray",
			args: args{
				nums: []int{10, 20, 30},
				k:    5,
			},
			want: 0,
		},
		{
			name: "all_elements_valid",
			args: args{
				nums: []int{1, 2, 3},
				k:    20,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
