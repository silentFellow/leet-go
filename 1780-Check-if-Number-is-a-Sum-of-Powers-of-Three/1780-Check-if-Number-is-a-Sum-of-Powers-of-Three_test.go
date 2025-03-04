package leetcode

import "testing"

func Test_checkPowersOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test case 1", args: args{n: 12}, want: true},
		{name: "Test case 2", args: args{n: 91}, want: true},
		{name: "Test case 3", args: args{n: 21}, want: false},
		{name: "Test case 4", args: args{n: 1}, want: true},
		{name: "Test case 5", args: args{n: 27}, want: true},
		{name: "Test case 6", args: args{n: 28}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPowersOfThree(tt.args.n); got != tt.want {
				t.Errorf("checkPowersOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
