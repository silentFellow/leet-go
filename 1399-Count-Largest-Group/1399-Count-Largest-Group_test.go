package leetcode

import "testing"

func Test_countLargestGroup(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test case 1", args: args{n: 13}, want: 4},
		{name: "Test case 2", args: args{n: 2}, want: 2},
		{name: "Test case 3", args: args{n: 15}, want: 6},
		{name: "Test case 4", args: args{n: 24}, want: 5},
		{name: "Test case 5", args: args{n: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLargestGroup(tt.args.n); got != tt.want {
				t.Errorf("countLargestGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfNo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfNo(tt.args.n); got != tt.want {
				t.Errorf("sumOfNo() = %v, want %v", got, tt.want)
			}
		})
	}
}
