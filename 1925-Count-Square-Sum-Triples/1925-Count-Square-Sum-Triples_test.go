package leetcode

import "testing"

func Test_countTriples(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n=1",
			args: args{n: 1},
			want: 0,
		},
		{
			name: "n=2",
			args: args{n: 2},
			want: 0,
		},
		{
			name: "n=5",
			args: args{n: 5},
			want: 2, // (3,4,5) and (4,3,5)
		},
		{
			name: "n=10",
			args: args{n: 10},
			want: 4, // (3,4,5), (4,3,5), (6,8,10), (8,6,10)
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriples(tt.args.n); got != tt.want {
				t.Errorf("countTriples() = %v, want %v", got, tt.want)
			}
		})
	}
}
