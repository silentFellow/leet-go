package leetcode

import "testing"

func Test_maximumInvitations(t *testing.T) {
	type args struct {
		favorite []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_case_1",
			args: args{favorite: []int{2, 2, 1, 2}},
			want: 3,
		},
		{
			name: "test_case_2",
			args: args{favorite: []int{1, 2, 0}},
			want: 3,
		},
		{
			name: "test_case_3",
			args: args{favorite: []int{3, 0, 1, 4, 1}},
			want: 4,
		},
		{
			name: "test_case_4",
			args: args{favorite: []int{1, 0}},
			want: 2,
		},
		{
			name: "test_case_5",
			args: args{favorite: []int{1, 2, 3, 4, 0}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumInvitations(tt.args.favorite); got != tt.want {
				t.Errorf("maximumInvitations() = %v, want %v", got, tt.want)
			}
		})
	}
}
