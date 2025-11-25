package leetcode

import "testing"

func Test_smallestRepunitDivByK(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"k=1", args{1}, 1},
		{"k=2", args{2}, -1},
		{"k=3", args{3}, 3},
		{"k=5", args{5}, -1},
		{"k=7", args{7}, 6},
		{"k=9", args{9}, 9},
		{"k=11", args{11}, 2},
		{"k=13", args{13}, 6},
		{"k=17", args{17}, 16},
		{"k=23", args{23}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRepunitDivByK(tt.args.k); got != tt.want {
				t.Errorf("smallestRepunitDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
