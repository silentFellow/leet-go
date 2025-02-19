package leetcode

import "testing"

func Test_getHappyString(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{n: 1, k: 1},
			want: "a",
		},
		{
			name: "Test case 2",
			args: args{n: 1, k: 2},
			want: "b",
		},
		{
			name: "Test case 3",
			args: args{n: 1, k: 3},
			want: "c",
		},
		{
			name: "Test case 4",
			args: args{n: 1, k: 4},
			want: "",
		},
		{
			name: "Test case 5",
			args: args{n: 3, k: 9},
			want: "cab",
		},
		{
			name: "Test case 6",
			args: args{n: 3, k: 12},
			want: "cbc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHappyString(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getHappyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
