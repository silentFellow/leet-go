package leetcode

import "testing"

func Test_countCoveredBuildings(t *testing.T) {
	type args struct {
		n         int
		buildings [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n:         3,
				buildings: [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}},
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				n:         3,
				buildings: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				n:         5,
				buildings: [][]int{{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3}},
			},
			want: 1,
		},
		{
			name: "No covered buildings, single row",
			args: args{
				n:         4,
				buildings: [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}},
			},
			want: 0,
		},
		{
			name: "No covered buildings, single column",
			args: args{
				n:         4,
				buildings: [][]int{{1, 1}, {2, 1}, {3, 1}, {4, 1}},
			},
			want: 0,
		},
		{
			name: "All buildings covered in a plus shape",
			args: args{
				n:         3,
				buildings: [][]int{{2, 1}, {2, 2}, {2, 3}, {1, 2}, {3, 2}},
			},
			want: 1,
		},
		{
			name: "Large grid, no covered",
			args: args{
				n:         100000,
				buildings: [][]int{{1, 1}, {100000, 100000}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCoveredBuildings(tt.args.n, tt.args.buildings); got != tt.want {
				t.Errorf("countCoveredBuildings() = %v, want %v", got, tt.want)
			}
		})
	}
}
