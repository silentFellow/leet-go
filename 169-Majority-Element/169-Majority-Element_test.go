package leetcode

import "testing"

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "single element",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "majority element in the middle",
			args: args{nums: []int{3, 2, 3}},
			want: 3,
		},
		{
			name: "majority element at the end",
			args: args{nums: []int{2, 2, 1, 1, 1, 2, 2}},
			want: 2,
		},
		{
			name: "all elements are the same",
			args: args{nums: []int{4, 4, 4, 4}},
			want: 4,
		},
		{
			name: "majority element at the start",
			args: args{nums: []int{5, 5, 5, 1, 2, 5, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
