package leetcode

import "testing"

func Test_findDifferentBinaryString(t *testing.T) {
	type args struct {
		nums []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{nums: []string{"01", "10"}},
			want: "00",
		},
		{
			name: "Test Case 2",
			args: args{nums: []string{"00", "01"}},
			want: "10",
		},
		{
			name: "Test Case 3",
			args: args{nums: []string{"111", "011", "001"}},
			want: "000",
		},
		{
			name: "Test Case 4",
			args: args{nums: []string{"0"}},
			want: "1",
		},
		{
			name: "Test Case 5",
			args: args{nums: []string{"1"}},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifferentBinaryString(tt.args.nums); got != tt.want {
				t.Errorf("findDifferentBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
