package leetcode

import "testing"

func Test_maxDiff(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{555}, 888},
		{"example2", args{9}, 8},
		{"leading_zero_case", args{10000}, 80000},
		{"all_same_digits", args{1111}, 8888},
		{"mixed_digits", args{9288}, 8700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDiff(tt.args.num); got != tt.want {
				t.Errorf("maxDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
