package leetcode

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 5,
		},
		{
			name: "example 2",
			args: args{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
		{
			name: "single day",
			args: args{prices: []int{5}},
			want: 0,
		},
		{
			name: "increasing prices",
			args: args{prices: []int{1, 2, 3, 4, 5}},
			want: 4,
		},
		{
			name: "decreasing prices",
			args: args{prices: []int{5, 4, 3, 2, 1}},
			want: 0,
		},
		{
			name: "mixed prices",
			args: args{prices: []int{3, 3, 5, 0, 0, 3, 1, 4}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
