package leetcode

import "testing"

func Test_fractionToDecimal(t *testing.T) {
	type args struct {
		numerator   int
		denominator int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {"simple fraction", args{1, 2}, "0.5"},
		// {"whole number", args{2, 1}, "2"},
		{"repeating decimal", args{4, 333}, "0.(012)"},
		// {"negative result", args{-50, 8}, "-6.25"},
		// {"zero numerator", args{0, -5}, "0"},
		// {"negative repeating", args{-1, -2147483648}, "0.0000000004656612873077392578125"},
		// {"simple repeating", args{1, 3}, "0.(3)"},
		// {"long repeating", args{1, 6}, "0.1(6)"},
		// {"no repeating", args{7, 8}, "0.875"},
		// {"large denominator", args{1, 2147483647}, "0.0000000004656612873077392572"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionToDecimal(tt.args.numerator, tt.args.denominator); got != tt.want {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
