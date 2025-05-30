package leetcode

import "testing"

func Test_numEquivDominoPairs(t *testing.T) {
	type args struct {
		dominoes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{dominoes: [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{dominoes: [][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}},
			want: 3,
		},
		{
			name: "No pairs",
			args: args{dominoes: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}},
			want: 0,
		},
		{
			name: "All identical",
			args: args{dominoes: [][]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}}},
			want: 6,
		},
		{
			name: "Mixed pairs",
			args: args{dominoes: [][]int{{1, 2}, {2, 1}, {1, 2}, {2, 1}}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numEquivDominoPairs(tt.args.dominoes); got != tt.want {
				t.Errorf("numEquivDominoPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
