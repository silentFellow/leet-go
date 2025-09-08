package leetcode

import "testing"

func Test_flowerGame(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"n=3,m=2", args{3, 2}, 3},
		{"n=1,m=1", args{1, 1}, 0},
		{"n=4,m=4", args{4, 4}, 8},
		{"n=12312,m=3", args{12312, 3}, 18468},
		{"n=42,m=24444", args{42, 24444}, 513324},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flowerGame(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("flowerGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
