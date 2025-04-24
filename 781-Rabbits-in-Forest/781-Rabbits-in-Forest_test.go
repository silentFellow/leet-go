package leetcode

import "testing"

func Test_numRabbits(t *testing.T) {
	type args struct {
		answers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{answers: []int{1, 1, 2}},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{answers: []int{10, 10, 10}},
			want: 11,
		},
		{
			name: "Single rabbit with no others",
			args: args{answers: []int{0}},
			want: 1,
		},
		{
			name: "All rabbits of the same color",
			args: args{answers: []int{3, 3, 3, 3}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRabbits(tt.args.answers); got != tt.want {
				t.Errorf("numRabbits() = %v, want %v", got, tt.want)
			}
		})
	}
}
