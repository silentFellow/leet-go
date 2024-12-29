package leetcode

import "testing"

func Test_repairCars(t *testing.T) {
	type args struct {
		ranks []int
		cars  int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example1",
			args: args{
				ranks: []int{4, 2, 3, 1},
				cars:  10,
			},
			want: 16,
		},
		{
			name: "example2",
			args: args{
				ranks: []int{5, 1, 8},
				cars:  6,
			},
			want: 16,
		},
		{
			name: "single mechanic",
			args: args{
				ranks: []int{3},
				cars:  5,
			},
			want: 75,
		},
		{
			name: "multiple mechanics, fewer cars",
			args: args{
				ranks: []int{2, 3, 1},
				cars:  3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repairCars(tt.args.ranks, tt.args.cars); got != tt.want {
				t.Errorf("repairCars() = %v, want %v", got, tt.want)
			}
		})
	}
}
