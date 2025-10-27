package leetcode

import "testing"

func Test_numberOfBeams(t *testing.T) {
	type args struct {
		bank []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{bank: []string{"011001", "000000", "010100", "001000"}},
			want: 8,
		},
		{
			name: "Example 2",
			args: args{bank: []string{"000", "111", "000"}},
			want: 0,
		},
		{
			name: "Single row",
			args: args{bank: []string{"10101"}},
			want: 0,
		},
		{
			name: "All empty rows",
			args: args{bank: []string{"000", "000", "000"}},
			want: 0,
		},
		{
			name: "Two rows with devices, no empty rows between",
			args: args{bank: []string{"101", "010"}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfBeams(tt.args.bank); got != tt.want {
				t.Errorf("numberOfBeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
