package leetcode

import "testing"

func Test_maximumUnits(t *testing.T) {
	type args struct {
		boxTypes  [][]int
		truckSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				boxTypes:  [][]int{{1, 3}, {2, 2}, {3, 1}},
				truckSize: 4,
			},
			want: 8,
		},
		{
			name: "example2",
			args: args{
				boxTypes:  [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},
				truckSize: 10,
			},
			want: 91,
		},
		{
			name: "edge_case1",
			args: args{
				boxTypes:  [][]int{{1, 1}},
				truckSize: 1,
			},
			want: 1,
		},
		{
			name: "edge_case2",
			args: args{
				boxTypes:  [][]int{{1, 1000}, {2, 999}, {3, 998}},
				truckSize: 2,
			},
			want: 1999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUnits(tt.args.boxTypes, tt.args.truckSize); got != tt.want {
				t.Errorf("maximumUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}
