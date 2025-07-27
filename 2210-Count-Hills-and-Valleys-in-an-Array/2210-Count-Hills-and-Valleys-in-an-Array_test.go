package leetcode

import "testing"

func Test_countHillValley(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "all equal",
			args: args{nums: []int{1, 1, 1, 1}},
			want: 0,
		},
		{
			name: "multiple hills and valleys",
			args: args{nums: []int{2, 4, 1, 1, 6, 5}},
			want: 3,
		},
		{
			name: "no hills or valleys",
			args: args{nums: []int{6, 6, 5, 5, 4, 1}},
			want: 0,
		},
		{
			name: "random case 2",
			args: args{nums: []int{2, 2, 1, 2, 2, 2}},
			want: 1,
		},
		{
			name: "random case 3",
			args: args{nums: []int{1, 2, 1, 1, 2, 2}},
			want: 2,
		},
		{
			name: "random case 4",
			args: args{nums: []int{2, 1, 1, 1, 2, 1}},
			want: 2,
		},
		{
			name: "random case 5",
			args: args{nums: []int{1, 2, 1, 2, 2, 1}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHillValley(tt.args.nums); got != tt.want {
				t.Errorf("countHillValley() = %v, want %v", got, tt.want)
			}
		})
	}
}
