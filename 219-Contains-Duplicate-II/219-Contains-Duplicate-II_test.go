package leetcode

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
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
			name: "example 1",
			args: args{nums: []int{1, 2, 3, 1}, k: 3},
			want: true,
		},
		{
			name: "example 2",
			args: args{nums: []int{1, 0, 1, 1}, k: 1},
			want: true,
		},
		{
			name: "example 3",
			args: args{nums: []int{1, 2, 3, 1, 2, 3}, k: 2},
			want: false,
		},
		{
			name: "no duplicates",
			args: args{nums: []int{1, 2, 3, 4, 5}, k: 3},
			want: false,
		},
		{
			name: "duplicates but not within k",
			args: args{nums: []int{1, 2, 3, 1, 2, 3}, k: 1},
			want: false,
		},
		{
			name: "duplicates within k",
			args: args{nums: []int{1, 2, 3, 4, 1}, k: 4},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
