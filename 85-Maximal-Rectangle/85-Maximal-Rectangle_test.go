package leetcode

import "testing"

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			}},
			want: 6,
		},
		{
			name: "Single zero",
			args: args{matrix: [][]byte{
				{'0'},
			}},
			want: 0,
		},
		{
			name: "Single one",
			args: args{matrix: [][]byte{
				{'1'},
			}},
			want: 1,
		},
		{
			name: "All ones",
			args: args{matrix: [][]byte{
				{'1', '1'},
				{'1', '1'},
			}},
			want: 4,
		},
		{
			name: "All zeros",
			args: args{matrix: [][]byte{
				{'0', '0'},
				{'0', '0'},
			}},
			want: 0,
		},
		{
			name: "Mixed small",
			args: args{matrix: [][]byte{
				{'1', '0', '1'},
				{'1', '1', '1'},
			}},
			want: 3,
		},
		{
			name: "Single row",
			args: args{matrix: [][]byte{
				{'1', '1', '0', '1'},
			}},
			want: 2,
		},
		{
			name: "Single column",
			args: args{matrix: [][]byte{
				{'1'},
				{'1'},
				{'0'},
				{'1'},
			}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
