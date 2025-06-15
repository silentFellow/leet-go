package leetcode

import "testing"

func Test_minMaxDifference(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{num: 11891},
			want: 99009,
		},
		{
			name: "example2",
			args: args{num: 90},
			want: 99,
		},
		{
			name: "single digit",
			args: args{num: 5},
			want: 9 - 0,
		},
		{
			name: "all same digits",
			args: args{num: 1111},
			want: 9999 - 0,
		},
		{
			name: "no remap needed",
			args: args{num: 909},
			want: 999 - 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMaxDifference(tt.args.num); got != tt.want {
				t.Errorf("minMaxDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
