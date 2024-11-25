package leetcode

import "testing"

func Test_kthLargestNumber(t *testing.T) {
	type args struct {
		nums []string
		k    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				nums: []string{"3", "6", "7", "10"},
				k:    4,
			},
			want: "3",
		},
		{
			name: "Example 2",
			args: args{
				nums: []string{"2", "21", "12", "1"},
				k:    3,
			},
			want: "2",
		},
		{
			name: "Example 3",
			args: args{
				nums: []string{"0", "0"},
				k:    2,
			},
			want: "0",
		},
		{
			name: "Single element",
			args: args{
				nums: []string{"5"},
				k:    1,
			},
			want: "5",
		},
		{
			name: "Large numbers",
			args: args{
				nums: []string{"100", "1000", "10", "10000"},
				k:    2,
			},
			want: "1000",
		},
		{
			name: "Same numbers",
			args: args{
				nums: []string{"1", "1", "1", "1"},
				k:    3,
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestNumber(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("kthLargestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
