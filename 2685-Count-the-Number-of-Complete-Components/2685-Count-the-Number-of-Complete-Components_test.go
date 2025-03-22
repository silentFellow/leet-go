package leetcode

import "testing"

func Test_countCompleteComponents(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}},
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}},
			},
			want: 1,
		},
		{
			name: "No edges",
			args: args{
				n:     4,
				edges: [][]int{},
			},
			want: 4,
		},
		{
			name: "Single complete component",
			args: args{
				n:     3,
				edges: [][]int{{0, 1}, {1, 2}, {0, 2}},
			},
			want: 1,
		},
		{
			name: "Disconnected nodes",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {2, 3}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCompleteComponents(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("countCompleteComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
