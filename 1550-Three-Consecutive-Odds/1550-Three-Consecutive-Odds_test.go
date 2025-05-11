package leetcode

import "testing"

func Test_threeConsecutiveOdds(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no odds",
			args: args{arr: []int{2, 6, 4, 1}},
			want: false,
		},
		{
			name: "has three consecutive odds",
			args: args{arr: []int{1, 2, 34, 3, 4, 5, 7, 23, 12}},
			want: true,
		},
		{
			name: "exactly three odds at start",
			args: args{arr: []int{1, 3, 5}},
			want: true,
		},
		{
			name: "less than three odds",
			args: args{arr: []int{1, 3}},
			want: false,
		},
		{
			name: "three odds in the middle",
			args: args{arr: []int{2, 1, 3, 5, 4}},
			want: true,
		},
		{
			name: "three odds at end",
			args: args{arr: []int{2, 4, 1, 3, 5}},
			want: true,
		},
		{
			name: "no three consecutive odds, scattered",
			args: args{arr: []int{1, 2, 3, 4, 5}},
			want: false,
		},
		{
			name: "all odds",
			args: args{arr: []int{1, 3, 5, 7, 9}},
			want: true,
		},
		{
			name: "single element",
			args: args{arr: []int{1}},
			want: false,
		},
		{
			name: "two elements",
			args: args{arr: []int{1, 3}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeConsecutiveOdds(tt.args.arr); got != tt.want {
				t.Errorf("threeConsecutiveOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}
