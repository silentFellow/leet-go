package leetcode

import "testing"

func Test_maximum69Number(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{9669}, 9969},
		{"example2", args{9996}, 9999},
		{"example3", args{9999}, 9999},
		{"single6", args{6}, 9},
		{"single9", args{9}, 9},
		{"all6s", args{6666}, 9666},
		{"all9s", args{999}, 999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum69Number(tt.args.num); got != tt.want {
				t.Errorf("maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
