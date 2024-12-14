package leetcode

import "testing"

func Test_numberOfCuts(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "Test 1", args{n: 4}, 2 },
		{ "Test 2", args{n: 3}, 3 },
		{ "Test 3", args{n: 1}, 0 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfCuts(tt.args.n); got != tt.want {
				t.Errorf("numberOfCuts() = %v, want %v", got, tt.want)
			}
		})
	}
}
