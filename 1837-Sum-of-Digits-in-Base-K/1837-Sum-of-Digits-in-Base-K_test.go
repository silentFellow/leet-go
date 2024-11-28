package leetcode

import "testing"

func Test_sumBase(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{n: 34, k: 6},
			want: 9,
		},
		{
			name: "Test case 2",
			args: args{n: 10, k: 10},
			want: 1,
		},
		{
			name: "Test case 3",
			args: args{n: 100, k: 2},
			want: 3,
		},
		{
			name: "Test case 4",
			args: args{n: 7, k: 3},
			want: 3,
		},
		{
			name: "Test case 5",
			args: args{n: 15, k: 5},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumBase(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("sumBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
