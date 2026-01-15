package leetcode

import "testing"

func Test_maximizeSquareHoleArea(t *testing.T) {
	type args struct {
		n     int
		m     int
		hBars []int
		vBars []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{n: 2, m: 1, hBars: []int{2, 3}, vBars: []int{2}},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{n: 1, m: 1, hBars: []int{2}, vBars: []int{2}},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{n: 2, m: 3, hBars: []int{2, 3}, vBars: []int{2, 4}},
			want: 4,
		},
		{
			name: "Non-consecutive bars",
			args: args{n: 5, m: 5, hBars: []int{2, 4}, vBars: []int{3, 5}},
			want: 4,
		},
		{
			name: "All consecutive bars",
			args: args{n: 4, m: 4, hBars: []int{2, 3, 4, 5}, vBars: []int{2, 3, 4, 5}},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeSquareHoleArea(tt.args.n, tt.args.m, tt.args.hBars, tt.args.vBars); got != tt.want {
				t.Errorf("maximizeSquareHoleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
