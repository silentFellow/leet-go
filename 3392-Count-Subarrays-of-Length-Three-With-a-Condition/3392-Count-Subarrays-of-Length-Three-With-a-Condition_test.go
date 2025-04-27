package leetcode

import "testing"

func Test_countSubarrays(t *testing.T) {
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
			args: args{nums: []int{1, 2, 1, 4, 1}},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 1, 1}},
			want: 0,
		},
		{
			name: "No valid subarrays",
			args: args{nums: []int{1, 3, 5, 7, 9}},
			want: 0,
		},
		{
			name: "Multiple valid subarrays",
			args: args{nums: []int{2, 8, 2, 4, 16, 4}},
			want: 2,
		},
		{
			name: "Negative numbers",
			args: args{nums: []int{-2, -8, -2, -4, -16, -4}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
