package leetcode

import "testing"

func Test_getDescentPeriods(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "single element",
			args: args{prices: []int{1}},
			want: 1,
		},
		{
			name: "all same",
			args: args{prices: []int{5, 5, 5}},
			want: 3,
		},
		{
			name: "strict descent",
			args: args{prices: []int{3, 2, 1}},
			want: 6,
		},
		{
			name: "mixed",
			args: args{prices: []int{3, 2, 1, 4}},
			want: 7,
		},
		{
			name: "no descent",
			args: args{prices: []int{8, 6, 7, 7}},
			want: 4,
		},
		{
			name: "longer descent",
			args: args{prices: []int{5, 4, 3, 2, 1}},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDescentPeriods(tt.args.prices); got != tt.want {
				t.Errorf("getDescentPeriods() = %v, want %v", got, tt.want)
			}
		})
	}
}
