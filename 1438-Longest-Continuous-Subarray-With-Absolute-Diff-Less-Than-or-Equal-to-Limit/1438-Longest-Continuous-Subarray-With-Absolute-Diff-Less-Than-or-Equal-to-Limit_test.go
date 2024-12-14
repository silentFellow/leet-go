package leetcode

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums:  []int{8, 2, 4, 7},
				limit: 4,
			},
			want: 2,
		},
		{
			name: "example2",
			args: args{
				nums:  []int{10, 1, 2, 4, 7, 2},
				limit: 5,
			},
			want: 4,
		},
		{
			name: "example3",
			args: args{
				nums:  []int{4, 2, 2, 2, 4, 4, 2, 2},
				limit: 0,
			},
			want: 3,
		},
		{
			name: "single_element",
			args: args{
				nums:  []int{1},
				limit: 0,
			},
			want: 1,
		},
		{
			name: "all_elements_same",
			args: args{
				nums:  []int{5, 5, 5, 5},
				limit: 0,
			},
			want: 4,
		},
		{
			name: "large_limit",
			args: args{
				nums:  []int{1, 3, 6, 10, 15},
				limit: 20,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
