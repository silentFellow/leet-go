package leetcode

import "testing"

func Test_lengthAfterTransformations(t *testing.T) {
	type args struct {
		s string
		t int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{s: "abcyy", t: 2},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{s: "azbk", t: 1},
			want: 5,
		},
		{
			name: "All z, t=1",
			args: args{s: "zzz", t: 1},
			want: 6,
		},
		{
			name: "Single a, t=3",
			args: args{s: "a", t: 3},
			want: 1,
		},
		{
			name: "Single z, t=2",
			args: args{s: "z", t: 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthAfterTransformations(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("lengthAfterTransformations() = %v, want %v", got, tt.want)
			}
		})
	}
}
