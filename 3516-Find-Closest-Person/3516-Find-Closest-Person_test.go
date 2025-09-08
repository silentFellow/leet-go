package leetcode

import "testing"

func Test_findClosest(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{2, 7, 4}, 1},
		{"Example 2", args{2, 5, 6}, 2},
		{"Example 3", args{1, 5, 3}, 0},
		{"Both at z", args{4, 4, 4}, 0},
		{"Person 1 closer", args{3, 10, 5}, 1},
		{"Person 2 closer", args{8, 2, 3}, 2},
		{"Negative positions, tie", args{-2, 8, 3}, 0},
		{"Negative positions, person 1 closer", args{-1, 10, 0}, 1},
		{"Negative positions, person 2 closer", args{-10, -2, -3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosest(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("findClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
