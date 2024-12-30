package leetcode

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Fibonacci of 0", args: args{n: 0}, want: 0},
		{name: "Fibonacci of 1", args: args{n: 1}, want: 1},
		{name: "Fibonacci of 2", args: args{n: 2}, want: 1},
		{name: "Fibonacci of 3", args: args{n: 3}, want: 2},
		{name: "Fibonacci of 4", args: args{n: 4}, want: 3},
		{name: "Fibonacci of 5", args: args{n: 5}, want: 5},
		{name: "Fibonacci of 6", args: args{n: 6}, want: 8},
		{name: "Fibonacci of 7", args: args{n: 7}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
