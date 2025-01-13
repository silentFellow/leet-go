package leetcode

import "testing"

func Test_areOccurrencesEqual(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{s: "abacbc"},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{s: "aaabb"},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{s: "aabbcc"},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{s: "abc"},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{s: "a"},
			want: true,
		},
		{
			name: "Test case 6",
			args: args{s: "aabbccc"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areOccurrencesEqual(tt.args.s); got != tt.want {
				t.Errorf("areOccurrencesEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
