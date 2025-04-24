package leetcode

import "testing"

func Test_countCompleteSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{1, 3, 1, 2, 2}},
			want: 4,
		},
		// {
		// 	name: "Example 2",
		// 	args: args{nums: []int{5, 5, 5, 5}},
		// 	want: 10,
		// },
		// {
		// 	name: "Single element",
		// 	args: args{nums: []int{1}},
		// 	want: 1,
		// },
		// {
		// 	name: "All distinct elements",
		// 	args: args{nums: []int{1, 2, 3, 4}},
		// 	want: 10,
		// },
		// {
		// 	name: "Mixed elements",
		// 	args: args{nums: []int{1, 2, 1, 3, 2}},
		// 	want: 5,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCompleteSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("countCompleteSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
