package leetcode

import "testing"

func Test_minCapability(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums: []int{2, 3, 5, 9},
				k:    2,
			},
			want: 5,
		},
		{
			name: "example2",
			args: args{
				nums: []int{2, 7, 9, 3, 1},
				k:    2,
			},
			want: 2,
		},
		{
			name: "single house",
			args: args{
				nums: []int{10},
				k:    1,
			},
			want: 10,
		},
		{
			name: "all houses same value",
			args: args{
				nums: []int{5, 5, 5, 5, 5},
				k:    3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCapability(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minCapability() = %v, want %v", got, tt.want)
			}
		})
	}
}
