package leetcode

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		h     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{piles: []int{3, 6, 7, 11}, h: 8},
			want: 4,
		},
		{
			name: "example 2",
			args: args{piles: []int{30, 11, 23, 4, 20}, h: 5},
			want: 30,
		},
		{
			name: "example 3",
			args: args{piles: []int{30, 11, 23, 4, 20}, h: 6},
			want: 23,
		},
		{
			name: "example 4",
			args: args{piles: []int{312884470}, h: 968709470},
			want: 1,
		},
		{
			name: "single pile",
			args: args{piles: []int{100}, h: 10},
			want: 10,
		},
		{
			name: "multiple piles, exact hours",
			args: args{piles: []int{3, 6, 7, 11}, h: 4},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.h); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
