package leetcode

import "testing"

func Test_kLengthApart(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{nums: []int{1, 0, 0, 0, 1, 0, 0, 1}, k: 2},
			want: true,
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 0, 0, 1, 0, 1}, k: 2},
			want: false,
		},
		{
			name: "No ones",
			args: args{nums: []int{0, 0, 0}, k: 2},
			want: true,
		},
		{
			name: "All ones, k=0",
			args: args{nums: []int{1, 1, 1, 1}, k: 0},
			want: true,
		},
		{
			name: "All ones, k=1",
			args: args{nums: []int{1, 1, 1, 1}, k: 1},
			want: false,
		},
		{
			name: "Single one",
			args: args{nums: []int{0, 0, 1, 0, 0}, k: 3},
			want: true,
		},
		{
			name: "Edge case, k larger than array",
			args: args{nums: []int{1, 0, 1}, k: 5},
			want: false,
		},
		{
			name: "Ones at start and end, enough distance",
			args: args{nums: []int{1, 0, 0, 0, 0, 1}, k: 4},
			want: true,
		},
		{
			name: "Ones at start and end, not enough distance",
			args: args{nums: []int{1, 0, 0, 1}, k: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kLengthApart(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("kLengthApart() = %v, want %v", got, tt.want)
			}
		})
	}
}
