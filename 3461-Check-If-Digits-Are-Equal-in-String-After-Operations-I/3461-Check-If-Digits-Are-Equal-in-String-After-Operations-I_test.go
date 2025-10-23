package leetcode

import "testing"

func Test_hasSameDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{s: "3902"},
			want: true,
		},
		// {
		// 	name: "Example 2",
		// 	args: args{s: "34789"},
		// 	want: false,
		// },
		// {
		// 	name: "All same digits",
		// 	args: args{s: "1111"},
		// 	want: true,
		// },
		// {
		// 	name: "Long string, true",
		// 	args: args{s: "4867716181911413979960155821878"},
		// 	want: false,
		// },
		// {
		// 	name: "Long string, false",
		// 	args: args{s: "472371665226339635424494604455723348326857778972516313101962164"},
		// 	want: false,
		// },
		// {
		// 	name: "Minimum length, true",
		// 	args: args{s: "111"},
		// 	want: true,
		// },
		// {
		// 	name: "Minimum length, false",
		// 	args: args{s: "123"},
		// 	want: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasSameDigits(tt.args.s); got != tt.want {
				t.Errorf("hasSameDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
