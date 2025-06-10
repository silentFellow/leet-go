package leetcode

import "testing"

func Test_maxFrequencyElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{nums: []int{1, 2, 2, 3, 1, 4}},
			want: 4,
		},
		{
			name: "example 2",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
		{
			name: "all same",
			args: args{nums: []int{7, 7, 7, 7}},
			want: 4,
		},
		{
			name: "two max freq elements",
			args: args{nums: []int{1, 2, 2, 3, 3}},
			want: 4,
		},
		{
			name: "single element",
			args: args{nums: []int{42}},
			want: 1,
		},
		{
			name: "empty input",
			args: args{nums: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequencyElements(tt.args.nums); got != tt.want {
				t.Errorf("maxFrequencyElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
