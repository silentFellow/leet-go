package leetcode

import "testing"

func Test_maxKelements(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test case 1",
			args: args{
				nums: []int{10, 10, 10, 10, 10},
				k:    5,
			},
			want: 50,
		},
		{
			name: "Test case 2",
			args: args{
				nums: []int{1, 10, 3, 3, 3},
				k:    3,
			},
			want: 17,
		},
		{
			name: "Test case 3",
			args: args{
				nums: []int{5, 5, 5, 5, 5},
				k:    3,
			},
			want: 15,
		},
		{
			name: "Test case 4",
			args: args{
				nums: []int{9, 8, 7, 6, 5},
				k:    2,
			},
			want: 17,
		},
		{
			name: "Test case 5",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxKelements(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxKelements() = %v, want %v", got, tt.want)
			}
		})
	}
}
