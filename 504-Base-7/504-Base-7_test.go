package leetcode

import "testing"

func Test_convertToBase7(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive number",
			args: args{num: 100},
			want: "202",
		},
		{
			name: "Negative number",
			args: args{num: -7},
			want: "-10",
		},
		{
			name: "Zero",
			args: args{num: 0},
			want: "0",
		},
		{
			name: "Small positive number",
			args: args{num: 5},
			want: "5",
		},
		{
			name: "Small negative number",
			args: args{num: -5},
			want: "-5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBase7(tt.args.num); got != tt.want {
				t.Errorf("convertToBase7() = %v, want %v", got, tt.want)
			}
		})
	}
}
