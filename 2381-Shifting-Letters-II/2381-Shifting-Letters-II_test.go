package leetcode

import "testing"

func Test_shiftingLetters(t *testing.T) {
	type args struct {
		s      string
		shifts [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{
				s:      "abc",
				shifts: [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}},
			},
			want: "ace",
		},
		{
			name: "example2",
			args: args{
				s:      "dztz",
				shifts: [][]int{{0, 0, 0}, {1, 1, 1}},
			},
			want: "catz",
		},
		{
			name: "no_shifts",
			args: args{
				s:      "abc",
				shifts: [][]int{},
			},
			want: "abc",
		},
		{
			name: "single_shift_forward",
			args: args{
				s:      "xyz",
				shifts: [][]int{{0, 2, 1}},
			},
			want: "yza",
		},
		{
			name: "single_shift_backward",
			args: args{
				s:      "abc",
				shifts: [][]int{{0, 2, 0}},
			},
			want: "zab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftingLetters(tt.args.s, tt.args.shifts); got != tt.want {
				t.Errorf("shiftingLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
