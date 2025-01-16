package leetcode

import "testing"

func Test_xorAllNums(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums1: []int{2, 1, 3},
				nums2: []int{10, 2, 5, 0},
			},
			want: 13,
		},
		{
			name: "example2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 0,
		},
		{
			name: "single_element",
			args: args{
				nums1: []int{1},
				nums2: []int{1},
			},
			want: 0,
		},
		{
			name: "large_numbers",
			args: args{
				nums1: []int{1000000000, 999999999},
				nums2: []int{999999998, 1000000000},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorAllNums(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("xorAllNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
