package leetcode

import "testing"

func Test_matchPlayersAndTrainers(t *testing.T) {
	type args struct {
		players  []int
		trainers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{players: []int{4, 7, 9}, trainers: []int{8, 2, 5, 8}},
			want: 2,
		},
		{
			name: "example 2",
			args: args{players: []int{1, 1, 1}, trainers: []int{10}},
			want: 1,
		},
		{
			name: "no match",
			args: args{players: []int{10, 20}, trainers: []int{1, 2}},
			want: 0,
		},
		{
			name: "all match",
			args: args{players: []int{1, 2, 3}, trainers: []int{3, 3, 3}},
			want: 3,
		},
		{
			name: "more trainers than players",
			args: args{players: []int{2, 4}, trainers: []int{1, 2, 3, 4, 5}},
			want: 2,
		},
		{
			name: "more players than trainers",
			args: args{players: []int{1, 2, 3, 4, 5}, trainers: []int{2, 4}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchPlayersAndTrainers(tt.args.players, tt.args.trainers); got != tt.want {
				t.Errorf("matchPlayersAndTrainers() = %v, want %v", got, tt.want)
			}
		})
	}
}
