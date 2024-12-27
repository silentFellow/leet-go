package leetcode

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		},
		{
			name: "example2",
			args: args{height: []int{4, 2, 0, 3, 2, 5}},
			want: 9,
		},
		{
			name: "no_trap",
			args: args{height: []int{1, 2, 3, 4, 5}},
			want: 0,
		},
		{
			name: "flat_surface",
			args: args{height: []int{0, 0, 0, 0}},
			want: 0,
		},
		{
			name: "single_peak",
			args: args{height: []int{0, 2, 0}},
			want: 0,
		},
		{
			name: "multiple_peaks",
			args: args{height: []int{3, 0, 2, 0, 4}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
