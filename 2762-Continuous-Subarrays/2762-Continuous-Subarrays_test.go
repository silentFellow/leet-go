package leetcode

import "testing"

func Test_continuousSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example1",
			args: args{nums: []int{5, 4, 2, 4}},
			want: 8,
		},
		{
			name: "example2",
			args: args{nums: []int{1, 2, 3}},
			want: 6,
		},
		{
			name: "single element",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "all elements same",
			args: args{nums: []int{2, 2, 2}},
			want: 6,
		},
		{
			name: "no continuous subarrays",
			args: args{nums: []int{1, 4, 7}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := continuousSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("continuousSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
