package leetcode

import "testing"

func Test_numTilings(t *testing.T) {
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
			want: 1,
		},
		{
			name: "n=2",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "n=3",
			args: args{n: 3},
			want: 5,
		},
		{
			name: "n=4",
			args: args{n: 4},
			want: 11,
		},
		{
			name: "n=5",
			args: args{n: 5},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTilings(tt.args.n); got != tt.want {
				t.Errorf("numTilings() = %v, want %v", got, tt.want)
			}
		})
	}
}
