package leetcode

import "testing"

func Test_canBeValid(t *testing.T) {
	type args struct {
		s      string
		locked string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{s: "))()))", locked: "010100"},
			want: true,
		},
		{
			name: "Example 2",
			args: args{s: "()()", locked: "0000"},
			want: true,
		},
		{
			name: "Example 3",
			args: args{s: ")", locked: "0"},
			want: false,
		},
		{
			name: "Example 4",
			args: args{s: ")(", locked: "00"},
			want: true,
		},
		{
			name: "All locked",
			args: args{s: "()", locked: "11"},
			want: true,
		},
		{
			name: "All unlocked",
			args: args{s: "(", locked: "0"},
			want: false,
		},
		{
			name: "Mixed locked and unlocked",
			args: args{s: "(()", locked: "101"},
			want: false,
		},
		{
			name: "Complex case",
			args: args{s: "((())())", locked: "110010011"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeValid(tt.args.s, tt.args.locked); got != tt.want {
				t.Errorf("canBeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
