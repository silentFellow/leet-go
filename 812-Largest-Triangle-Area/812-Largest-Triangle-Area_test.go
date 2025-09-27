package leetcode

import "testing"

func Test_largestTriangleArea(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example 1",
			args: args{points: [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}},
			want: 2.0,
		},
		{
			name: "Example 2",
			args: args{points: [][]int{{1, 0}, {0, 0}, {0, 1}}},
			want: 0.5,
		},
		{
			name: "Colinear points",
			args: args{points: [][]int{{0, 0}, {1, 1}, {2, 2}}},
			want: 0.0,
		},
		{
			name: "Triangle with negative coordinates",
			args: args{points: [][]int{{-1, -1}, {-1, 1}, {1, 1}}},
			want: 2.0,
		},
		{
			name: "All points on X axis",
			args: args{points: [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
			want: 0.0,
		},
		{
			name: "Large triangle",
			args: args{points: [][]int{{-50, -50}, {50, -50}, {0, 50}}},
			want: 5000.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestTriangleArea(tt.args.points); got != tt.want {
				t.Errorf("largestTriangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
