package leetcode

import "testing"

func Test_findMaxForm(t *testing.T) {
	type args struct {
		strs []string
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{strs: []string{"10", "0001", "111001", "1", "0"}, m: 5, n: 3},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{strs: []string{"10", "0", "1"}, m: 1, n: 1},
			want: 2,
		},
		{
			name: "Edge case: all fit",
			args: args{strs: []string{"0", "1"}, m: 1, n: 1},
			want: 2,
		},
		{
			name: "Edge case: none fit",
			args: args{strs: []string{"111", "000"}, m: 0, n: 0},
			want: 0,
		},
		{
			name: "Single string fits",
			args: args{strs: []string{"01"}, m: 1, n: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForm(tt.args.strs, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMaxForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
