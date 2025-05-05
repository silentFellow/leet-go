package leetcode

import "testing"

func Test_pushDominoes(t *testing.T) {
	type args struct {
		dominoes string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{dominoes: "RR.L"},
			want: "RR.L",
		},
		{
			name: "Example 2",
			args: args{dominoes: ".L.R...LR..L.."},
			want: "LL.RR.LLRRLL..",
		},
		{
			name: "All standing",
			args: args{dominoes: "....."},
			want: ".....",
		},
		{
			name: "All falling right",
			args: args{dominoes: "R...."},
			want: "RRRRR",
		},
		{
			name: "All falling left",
			args: args{dominoes: "....L"},
			want: "LLLLL",
		},
		{
			name: "Mixed forces",
			args: args{dominoes: "R...L"},
			want: "RR.LL",
		},
		{
			name: "Single domino",
			args: args{dominoes: "."},
			want: ".",
		},
		{
			name: "Alternating forces",
			args: args{dominoes: "R.L"},
			want: "R.L",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushDominoes(tt.args.dominoes); got != tt.want {
				t.Errorf("pushDominoes() = %v, want %v", got, tt.want)
			}
		})
	}
}
