package leetcode

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{2736}, 7236},
		{"example2", args{9973}, 9973},
		{"single digit", args{5}, 5},
		{"all same digits", args{1111}, 1111},
		{"leading zero after swap", args{98368}, 98863},
		{"descending order", args{98765432}, 98765432},
		{"swap at end", args{1993}, 9913},
		{"multiple same max digits", args{98800435}, 98850430},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
