package leetcode

import "testing"

func Test_countPoints(t *testing.T) {
	type args struct {
		rings string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{rings: "B0B6G0R6R0R6G9"},
			want: 1,
		},
		{
			name: "example 2",
			args: args{rings: "B0R0G0R9R0B0G0"},
			want: 1,
		},
		{
			name: "example 3",
			args: args{rings: "G4"},
			want: 0,
		},
		{
			name: "all rods with all colors",
			args: args{rings: "R0G0B0R1G1B1R2G2B2R3G3B3R4G4B4R5G5B5R6G6B6R7G7B7R8G8B8R9G9B9"},
			want: 10,
		},
		{
			name: "no rods with all colors",
			args: args{rings: "R0R1R2R3R4R5R6R7R8R9"},
			want: 0,
		},
		{
			name: "multiple rods with all colors",
			args: args{rings: "R0G0B0R1G1B1R2G2B2"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPoints(tt.args.rings); got != tt.want {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
