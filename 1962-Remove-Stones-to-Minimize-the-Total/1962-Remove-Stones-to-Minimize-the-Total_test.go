package leetcode

import "testing"

func Test_minStoneSum(t *testing.T) {
	type args struct {
		piles []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				piles: []int{5, 4, 9},
				k:     2,
			},
			want: 12,
		},
		{
			name: "example2",
			args: args{
				piles: []int{4, 3, 6, 7},
				k:     3,
			},
			want: 12,
		},
		{
			name: "no operations",
			args: args{
				piles: []int{1, 2, 3},
				k:     0,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStoneSum(tt.args.piles, tt.args.k); got != tt.want {
				t.Errorf("minStoneSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
