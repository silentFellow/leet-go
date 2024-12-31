package leetcode

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{s: "PAYPALISHIRING", numRows: 3},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "Example 2",
			args: args{s: "PAYPALISHIRING", numRows: 4},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "Example 3",
			args: args{s: "A", numRows: 1},
			want: "A",
		},
		{
			name: "Single row",
			args: args{s: "HELLO", numRows: 1},
			want: "HELLO",
		},
		{
			name: "Two rows",
			args: args{s: "HELLO", numRows: 2},
			want: "HLOEL",
		},
		{
			name: "Empty string",
			args: args{s: "", numRows: 3},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
