package leetcode

import "testing"

func Test_findNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{12, 345, 2, 6, 7896}},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{nums: []int{555, 901, 482, 1771}},
			want: 1,
		},
		{
			name: "All odd digits",
			args: args{nums: []int{1, 3, 5, 7, 9}},
			want: 0,
		},
		{
			name: "All even digits",
			args: args{nums: []int{22, 44, 66, 88}},
			want: 4,
		},
		{
			name: "Mixed digits",
			args: args{nums: []int{10, 123, 4567, 89012}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.args.nums); got != tt.want {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSize(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSize(tt.args.v); got != tt.want {
				t.Errorf("getSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
