package leetcode

import "testing"

func Test_maximumTripletValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{nums: []int{12, 6, 1, 2, 7}},
			want: 77,
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 10, 3, 4, 19}},
			want: 133,
		},
		{
			name: "Example 3",
			args: args{nums: []int{1, 2, 3}},
			want: 0,
		},
		{
			name: "Single peak",
			args: args{nums: []int{5, 1, 10, 1, 5}},
			want: 45,
		},
		{
			name: "All equal elements",
			args: args{nums: []int{5, 5, 5, 5, 5}},
			want: 0,
		},
		{
			name: "Large values",
			args: args{nums: []int{1000000, 1, 1000000}},
			want: 999999000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTripletValue(tt.args.nums); got != tt.want {
				t.Errorf("maximumTripletValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
