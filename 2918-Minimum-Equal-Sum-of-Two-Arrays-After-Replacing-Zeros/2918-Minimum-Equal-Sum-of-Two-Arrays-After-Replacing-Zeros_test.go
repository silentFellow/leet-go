package leetcode

import "testing"

func Test_minSum(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example1",
			args: args{
				nums1: []int{3, 2, 0, 1, 0},
				nums2: []int{6, 5, 0},
			},
			want: 12,
		},
		{
			name: "example2",
			args: args{
				nums1: []int{2, 0, 2, 0},
				nums2: []int{1, 4},
			},
			want: -1,
		},
		{
			name: "example3",
			args: args{
				nums1: []int{0, 0, 0},
				nums2: []int{0, 0, 0},
			},
			want: 3,
		},
		{
			name: "example4",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{4, 5, 6},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSum(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
