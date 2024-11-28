package leetcode

import "testing"

func Test_checkOnesSegment(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{s: "1001"},
			want: false,
		},
		{
			name: "Test case 2",
			args: args{s: "110"},
			want: true,
		},
		{
			name: "Test case 3",
			args: args{s: "1"},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{s: "0"},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{s: "1111"},
			want: true,
		},
		{
			name: "Test case 6",
			args: args{s: "101"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkOnesSegment(tt.args.s); got != tt.want {
				t.Errorf("checkOnesSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}
