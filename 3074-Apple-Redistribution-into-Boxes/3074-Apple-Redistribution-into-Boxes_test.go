package leetcode

import "testing"

func Test_minimumBoxes(t *testing.T) {
	type args struct {
		apple    []int
		capacity []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				apple:    []int{1, 3, 2},
				capacity: []int{4, 3, 1, 5, 2},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				apple:    []int{5, 5, 5},
				capacity: []int{2, 4, 2, 7},
			},
			want: 4,
		},
		{
			name: "All apples fit in one box",
			args: args{
				apple:    []int{10},
				capacity: []int{10, 1, 2, 3},
			},
			want: 1,
		},
		{
			name: "Each apple needs a separate box",
			args: args{
				apple:    []int{1, 1, 1},
				capacity: []int{1, 1, 1},
			},
			want: 3,
		},
		{
			name: "Large capacities, fewer boxes needed",
			args: args{
				apple:    []int{2, 2, 2, 2},
				capacity: []int{8, 1, 1, 1},
			},
			want: 1,
		},
		{
			name: "Minimum input size",
			args: args{
				apple:    []int{1},
				capacity: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumBoxes(tt.args.apple, tt.args.capacity); got != tt.want {
				t.Errorf("minimumBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
