package leetcode

import "testing"

func Test_numOfSubarrays(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{arr: []int{1, 3, 5}},
			want: 4,
		},
		{
			name: "example2",
			args: args{arr: []int{2, 4, 6}},
			want: 0,
		},
		{
			name: "example3",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7}},
			want: 16,
		},
		{
			name: "single element odd",
			args: args{arr: []int{1}},
			want: 1,
		},
		{
			name: "single element even",
			args: args{arr: []int{2}},
			want: 0,
		},
		{
			name: "mixed elements",
			args: args{arr: []int{1, 2, 3, 4, 5}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfSubarrays(tt.args.arr); got != tt.want {
				t.Errorf("numOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
