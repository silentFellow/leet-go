package leetcode

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example1",
			args: args{
				nums: []int{1, 3, 2, 3, 3},
				k:    2,
			},
			want: 6,
		},
		{
			name: "example2",
			args: args{
				nums: []int{1, 4, 2, 1},
				k:    3,
			},
			want: 0,
		},
		{
			name: "single_element",
			args: args{
				nums: []int{5},
				k:    1,
			},
			want: 1,
		},
		{
			name: "all_same_elements",
			args: args{
				nums: []int{2, 2, 2, 2},
				k:    3,
			},
			want: 3,
		},
		{
			name: "no_valid_subarray",
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    5,
			},
			want: 0,
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
