package leetcode

import "testing"

func Test_countOperations(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{2, 3}, 3},
		{"example2", args{10, 10}, 1},
		{"zero1", args{0, 5}, 0},
		{"zero2", args{7, 0}, 0},
		{"one_and_two", args{1, 2}, 2},
		{"large_diff", args{100000, 1}, 100000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOperations(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("countOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
