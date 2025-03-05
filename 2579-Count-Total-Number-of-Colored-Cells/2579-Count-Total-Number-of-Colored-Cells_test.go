package leetcode

import "testing"

func Test_coloredCells(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test case 1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "test case 2",
			args: args{n: 2},
			want: 5,
		},
		{
			name: "test case 3",
			args: args{n: 3},
			want: 13,
		},
		{
			name: "test case 4",
			args: args{n: 4},
			want: 25,
		},
		{
			name: "test case 5",
			args: args{n: 5},
			want: 41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coloredCells(tt.args.n); got != tt.want {
				t.Errorf("coloredCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
