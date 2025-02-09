package leetcode

import "testing"

func Test_countBadPairs(t *testing.T) {
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
			args: args{nums: []int{4, 1, 3, 3}},
			want: 5,
		},
		{
			name: "example2",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 0,
		},
		{
			name: "single_element",
			args: args{nums: []int{1}},
			want: 0,
		},
		{
			name: "all_same_elements",
			args: args{nums: []int{2, 2, 2, 2}},
			want: 6,
		},
		{
			name: "all_bad_pairs",
			args: args{nums: []int{5, 4, 3, 2, 1}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBadPairs(tt.args.nums); got != tt.want {
				t.Errorf("countBadPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
