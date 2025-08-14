package leetcode

import "testing"

func Test_largestGoodInteger(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example1", args{"6777133339"}, "777"},
		{"example2", args{"2300019"}, "000"},
		{"example3", args{"42352338"}, ""},
		{"example4", args{"99911122"}, "999"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestGoodInteger(tt.args.num); got != tt.want {
				t.Errorf("largestGoodInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
