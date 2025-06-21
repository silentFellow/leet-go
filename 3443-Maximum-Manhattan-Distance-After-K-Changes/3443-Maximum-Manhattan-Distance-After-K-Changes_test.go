package leetcode

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{s: "NWSE", k: 1},
			want: 3,
		},
		{
			name: "example 2",
			args: args{s: "NSWWEW", k: 3},
			want: 6,
		},
		{
			name: "no change, straight north",
			args: args{s: "NNNN", k: 0},
			want: 4,
		},
		{
			name: "all south, change all to north",
			args: args{s: "SSSS", k: 4},
			want: 4,
		},
		{
			name: "zigzag, no change",
			args: args{s: "NESW", k: 0},
			want: 2,
		},
		{
			name: "all west, no change",
			args: args{s: "WWWW", k: 0},
			want: 4,
		},
		{
			name: "empty string",
			args: args{s: "", k: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
