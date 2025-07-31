package leetcode

import "testing"

func Test_subarrayBitwiseORs(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Single zero",
			args: args{arr: []int{0}},
			want: 1,
		},
		{
			name: "All ones",
			args: args{arr: []int{1, 1, 1}},
			want: 1,
		},
		{
			name: "Example case 1",
			args: args{arr: []int{1, 1, 2}},
			want: 3,
		},
		{
			name: "Example case 2",
			args: args{arr: []int{1, 2, 4}},
			want: 6,
		},
		{
			name: "Increasing numbers",
			args: args{arr: []int{1, 2, 3, 4}},
			want: 5,
		},
		{
			name: "Zeros and ones",
			args: args{arr: []int{0, 1, 0, 1}},
			want: 2,
		},
		{
			name: "Longer array, same elements",
			args: args{arr: []int{5, 5, 5, 5}},
			want: 1,
		},
		{
			name: "Longer array, distinct",
			args: args{arr: []int{1, 2, 4, 8, 16}},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarrayBitwiseORs(tt.args.arr); got != tt.want {
				t.Errorf("subarrayBitwiseORs() = %v, want %v", got, tt.want)
			}
		})
	}
}
