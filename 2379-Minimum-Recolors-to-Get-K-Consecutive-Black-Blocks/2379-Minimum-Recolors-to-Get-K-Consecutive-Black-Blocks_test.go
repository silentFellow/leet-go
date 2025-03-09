package leetcode

import "testing"

func Test_minimumRecolors(t *testing.T) {
	type args struct {
		blocks string
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "example1",
		// 	args: args{
		// 		blocks: "WBBWWBBWBW",
		// 		k:      7,
		// 	},
		// 	want: 3,
		// },
		{
			name: "example2",
			args: args{
				blocks: "WBWBBBW",
				k:      2,
			},
			want: 0,
		},
		// {
		// 	name: "test1",
		// 	args: args{
		// 		blocks: "WWWWWW",
		// 		k:      3,
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "test2",
		// 	args: args{
		// 		blocks: "BBBBBB",
		// 		k:      4,
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "test3",
		// 	args: args{
		// 		blocks: "BWBWBWBWBW",
		// 		k:      5,
		// 	},
		// 	want: 2,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRecolors(tt.args.blocks, tt.args.k); got != tt.want {
				t.Errorf("minimumRecolors() = %v, want %v", got, tt.want)
			}
		})
	}
}
