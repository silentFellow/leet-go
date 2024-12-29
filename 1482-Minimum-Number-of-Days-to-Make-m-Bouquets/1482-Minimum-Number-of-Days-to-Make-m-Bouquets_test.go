package leetcode

import "testing"

func Test_minDays(t *testing.T) {
	type args struct {
		bloomDay []int
		m        int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_case_1",
			args: args{
				bloomDay: []int{1, 10, 3, 10, 2},
				m:        3,
				k:        1,
			},
			want: 3,
		},
		{
			name: "test_case_2",
			args: args{
				bloomDay: []int{1, 10, 3, 10, 2},
				m:        3,
				k:        2,
			},
			want: -1,
		},
		{
			name: "test_case_3",
			args: args{
				bloomDay: []int{7, 7, 7, 7, 12, 7, 7},
				m:        2,
				k:        3,
			},
			want: 12,
		},
		{
			name: "test_case_4",
			args: args{
				bloomDay: []int{1000000000, 1000000000},
				m:        1,
				k:        1,
			},
			want: 1000000000,
		},
		{
			name: "test_case_5",
			args: args{
				bloomDay: []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6},
				m:        4,
				k:        2,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays(tt.args.bloomDay, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("minDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
