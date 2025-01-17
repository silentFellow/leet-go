package leetcode

import "testing"

func Test_doesValidArrayExist(t *testing.T) {
	type args struct {
		derived []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{derived: []int{1, 1, 0}},
			want: true,
		},
		{
			name: "example2",
			args: args{derived: []int{1, 1}},
			want: true,
		},
		{
			name: "example3",
			args: args{derived: []int{1, 0}},
			want: false,
		},
		{
			name: "all zeros",
			args: args{derived: []int{0, 0, 0}},
			want: true,
		},
		{
			name: "all ones",
			args: args{derived: []int{1, 1, 1}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doesValidArrayExist(tt.args.derived); got != tt.want {
				t.Errorf("doesValidArrayExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
