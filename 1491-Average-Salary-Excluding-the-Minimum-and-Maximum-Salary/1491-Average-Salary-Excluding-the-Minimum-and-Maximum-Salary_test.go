package leetcode

import "testing"

func Test_average(t *testing.T) {
	type args struct {
		salary []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "example1",
			args: args{salary: []int{4000, 3000, 1000, 2000}},
			want: 2500.00000,
		},
		{
			name: "example2",
			args: args{salary: []int{1000, 2000, 3000}},
			want: 2000.00000,
		},
		{
			name: "example3",
			args: args{salary: []int{6000, 5000, 4000, 3000, 2000, 1000}},
			want: 3500.00000,
		},
		{
			name: "example4",
			args: args{salary: []int{8000, 7000, 6000, 5000, 4000, 3000, 2000, 1000}},
			want: 4500.00000,
		},
		{
			name: "example5",
			args: args{salary: []int{5000, 3000, 1000}},
			want: 3000.00000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.args.salary); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
