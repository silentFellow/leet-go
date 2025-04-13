package leetcode

import "testing"

func Test_countSymmetricIntegers(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "range_1_to_100",
			args: args{low: 1, high: 100},
			want: 9,
		},
		{
			name: "range_1200_to_1230",
			args: args{low: 1200, high: 1230},
			want: 4,
		},
		{
			name: "range_11_to_11",
			args: args{low: 11, high: 11},
			want: 1,
		},
		{
			name: "range_1_to_10",
			args: args{low: 1, high: 10},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSymmetricIntegers(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("countSymmetricIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
